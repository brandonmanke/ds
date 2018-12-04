/*Super basic architecture overview
 *       ----------------
 *       |  Publisher   |
 *       ----------------
 *              | (publish)
 *              v
 *          ---------
 *          | Redis |
 *          ---------
 *            |   ^
 * (Receive)  v   | (Post messages & Subscribe)
 *       ----------------
 *       |  Subscriber  |
 *       ----------------
 *               ^
 *               | (WebSockets)
 *               v
 *        --------------
 *        |   Client   |
 *        --------------
 */
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

func testMessages(pub *Publisher) {
	err := pub.PublishMessages("Hello", "World", "!!!!", "more messages", "omg")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//time.Sleep(500 * time.Millisecond)
}

// Reads config.json file at file path and returns it as a map
func readConfig(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config map[string]string
	json.Unmarshal(bytes, &config)
	return config, nil
}

func setInterval(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func parseInput(config map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your command: ")
	for text, err := reader.ReadString('\n'); err != nil || text != "quit"; {
		trimmed := strings.Trim(text, " \n")
		tokens := strings.Split(trimmed, " ")
		if len(tokens) != 2 || len(tokens) != 3 || tokens[0] != "pub" {
			//fmt.Println("Please enter a correct command")
			//fmt.Println("pub channelName [message/category]")
			continue
		}
		switch tokens[1] {
		case "weather":
			publishWeather(tokens[1], config)
		case "news":
			publishNews(tokens[1], config)
		case "test.channel":
			var sb strings.Builder
			for i := 2; i < len(tokens); i++ {
				sb.WriteString(tokens[i])
			}
			publishTest(tokens[1], sb.String())
		default:
			fmt.Println("Unrecognized input. Please try again.")
			break
		}
	}
}

func publishWeather(channel string, config map[string]string) error {
	pub, err := CreatePublisher(channel)
	weatherAPI := CreateWeatherAPI(config["weather-api-key"])
	res, err := weatherAPI.GetForecast()
	if err != nil {
		return err
	}

	serializedWeatherRes, err := SerializeJSON(res)
	if err != nil {
		return err
	}
	pub.PublishMessage(serializedWeatherRes)
	return nil
}

func publishNews(channel string, config map[string]string) error {
	pub, err := CreatePublisher(channel)
	newsAPI := CreateNYTimesAPI(config["news-api-key"])
	res, err := newsAPI.GetHome()
	if err != nil {
		return err
	}

	serializedNewsRes, err := SerializeJSON(res)
	if err != nil {
		return err
	}

	pub.PublishMessage(serializedNewsRes)
	return nil
}

func publishTest(channel, msg string) error {
	pub, err := CreatePublisher(channel)
	if err != nil {
		return err
	}
	pub.PublishMessage(msg)
	return nil
}

func main() {
	pub, err := CreatePublisher("test.channel")
	if err != nil {
		panic(err)
	}
	defer pub.Close()

	config, err := readConfig("./config.json")
	if err != nil {
		panic(err)
	}

	// TODO remove wait groups and use channels instead
	// to know when threads finish publishing
	// We need to refactor this anyways since we are going
	// to be polling a bunch of different APIs
	//ch := make

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		testMessages(pub)
		wg.Done()
	}()

	go func() {
		testMessages(pub)
		wg.Done()
	}()

	setInterval(60*time.Second, func() {
		publishTest("test.channel", "every 60 seconds")
	})

	wg.Wait()

	parseInput(config)
}
