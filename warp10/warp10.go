package warp10

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"bytes"
)

const (
	Name    = "warp10"
	Version = 1
	timeout = 30
)

type Warp10Publisher struct {
}

func (f *Warp10Publisher) Publish(metrics []plugin.Metric, cfg plugin.Config) error {
	logger := getLogger(cfg)
	logger.Debug("Publishing started")

	logger.Debug("publishing %v metrics to %v", len(metrics), cfg)
	warp_url, err := cfg.GetString("warp_url")
	if err != nil {
		return err
	}
	token, err := cfg.GetString("token")
	if err != nil {
		return err
	}

	var temp GTS
	pts := make([]string, 0)
	// Parsing
	for _, m := range metrics {
		ns := m.Namespace.Strings()

		tempTags := make(map[string]string)

		tags := m.Tags
		for k, v := range tags {
			tempTags[k] = string(v)
		}
		//tempTags["host"] = string(tags[core.STD_TAG_PLUGIN_RUNNING_ON])

		newtags := map[string]string{}
		isDynamic, indexes := m.Namespace.IsDynamic()
		if isDynamic {
			for i, j := range indexes {
				// The second return value from IsDynamic(), in this case `indexes`, is the index of
				// the dynamic element in the unmodified namespace. However, here we're deleting
				// elements, which is problematic when the number of dynamic elements in a namespace is
				// greater than 1. Therefore, we subtract i (the loop iteration) from j
				// (the original index) to compensate.
				//
				// Remove "data" from the namespace and create a tag for it
				ns = append(ns[:j-i], ns[j-i+1:]...)
				newtags[m.Namespace[j].Name] = m.Namespace[j].Value
			}
		}

		for k, v := range newtags {
			tempTags[k] = v
		}
		metricValue := fmt.Sprintf("%v",m.Data)

		tagsSlice := buildTags(tempTags)
		finalTags := strings.Join(tagsSlice, ",")

		temp = GTS{
			Timestamp: m.Timestamp.Unix()*1000000,
			Metric:strings.Join(ns, "."),
			Tags:finalTags,
			Value:metricValue,
		}
		messageLine := fmt.Sprintf("%d// %s{%s} %s\n", temp.Timestamp, temp.Metric, temp.Tags, temp.Value)
		logger.Debug("Metric ready to send %v",messageLine)
		pts = append(pts, messageLine)
	}

	payload := fmt.Sprint(strings.Join(pts, "\n"))
	req, err := http.NewRequest("POST", warp_url, bytes.NewBufferString(payload))
	req.Header.Set("X-CityzenData-Token", token)
	req.Header.Set("Content-Type", "text/plain")

	start := time.Now()
	client := &http.Client{Timeout:timeout * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Unable to send metrics. Error: %s", err)
		return fmt.Errorf("Unable to send metrics. Error: %s", err)
	}

	logger.Debug("Warp10 POST http code: %v",resp.StatusCode)
	elapsed := time.Since(start)
	logger.Debug("Warp10 POST took %v", elapsed)
	defer resp.Body.Close()
	logger.Debug("Metrics sent to Warp10.")
	return nil
}

func (f *Warp10Publisher) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	policy.AddNewStringRule([]string{""}, "warp_url", true)
	policy.AddNewStringRule([]string{""}, "token", true)
	policy.AddNewStringRule([]string{""}, "log-level", false)

	return *policy, nil
}

func getLogger(cfg plugin.Config) *log.Entry {
	logger := log.WithFields(log.Fields{
		"plugin-name":    Name,
		"plugin-version": Version,
		"plugin-type":    "publisher",
	})

	log.SetLevel(log.WarnLevel)

	levelValue, err := cfg.GetString("log-level")
	if err == nil {
		if level, err := log.ParseLevel(strings.ToLower(levelValue)); err == nil {
			log.SetLevel(level)
		} else {
			log.WithFields(log.Fields{
				"value":             strings.ToLower(levelValue),
				"acceptable values": "warn, error, debug, info",
			}).Warn("Invalid config value")
		}
	}
	return logger
}

func buildTags(ptTags map[string]string) []string {
	sizeTags := len(ptTags)
	//sizeTags += 1
	tags := make([]string, sizeTags)
	index := 0
	for k, v := range ptTags {
		tags[index] = fmt.Sprintf("%s=%s", k, v)
		index += 1
	}
	sort.Strings(tags)
	return tags
}
