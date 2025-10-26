package collector

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type CollectorConfig struct {
	LogFilePath  string
	PollInterval time.Duration
}

type Collector struct {
	config CollectorConfig
	stop   chan bool
}

// create new instance
func NewCollector(config CollectorConfig) *Collector {
	return &Collector{
		config: config,
		stop:   make(chan bool),
	}
}

func (c *Collector) readLog(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("[Collector] ERROR OPENING FILE: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// tODO: kirim ke storage
		fmt.Printf("[Collector] %s\n", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("[Collector] ERROR READING FILE: %v\n", err)
	}
}

func (c *Collector) Stop() {
	c.stop <- true
}

func (c *Collector) Start() {
	fmt.Println("[Collector] STARTING LOG COLLECTOR")
	filePath := c.config.LogFilePath
	for {
		select {
		case <-c.stop:
			fmt.Println("[Collector] STOPPING LOG COLLECTOR")
			return
		default:
			c.readLog(filePath)
			time.Sleep(c.config.PollInterval)
		}
	}
}
