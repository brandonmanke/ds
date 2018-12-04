import React, { Component } from "react";
import "./App.css";
import "./index.css";

import { Table, Grid, Row, Col, PageHeader } from "react-bootstrap";
import Feed from "./components/Feed";
import Subscriptions from "./components/Subscriptions";

import * as newsResponse from './news_response.json';
// import * as weatherResponse from "./weather_response.json";

console.log(newsResponse.results[0].abstract)

const newsFeed = newsResponse.results.map( result => ({
  title: result.abstract,
  time: new Date().toLocaleTimeString()
}))

console.log(newsFeed);

//Establishing socket instance
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onclose = evt => {
  console.log("CLOSE");
};

socket.onerror = evt => {
  console.log("ERROR: " + evt.data);
};

// let counter = 0;

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      // Dummy feed data. Will switch to null once ready to handle socket messages
  
      // feed: newsFeed,
  
      feed: '',
  
      // Array of state objects that change based on subscriptions
      subscriptions: [
        {
          title: "test.channel",
          subscribed: true
        },
        {
          title: "weather",
          subscribed: false
        },
        {
          title: "news",
          subscribed: false
        }
      ]
    };
    this.handleSubscription = this.handleSubscription.bind(this)
    this.handleUnsubscription = this.handleUnsubscription.bind(this)

  }
  
  // Might need to use this. Placeholder for now
  componentDidMount() {}

  //Handles button controls to subscribe to channel
  handleSubscription = button => {
    console.log(this.state.subscriptions[button].title, " subscribed")
    socket.send("subscribe " + this.state.subscriptions[button].title );
    this.setState(prevState => {
      return {
        // feed: [
        //   ...prevState.feed,
        //   {
        //     title: newsFeed[counter++].title,
        //     time: new Date().toLocaleTimeString()
        //   }
        // ],
        subscribed: (prevState.subscriptions[button].subscribed = true)
      };
    });
  };

  //Handles button for unsubscribe
  handleUnsubscription = button => {
    console.log(this.state.subscriptions[button].title, " unsubscribed")
    socket.send("unsubscribe " + this.state.subscriptions[button].title );
    this.setState(prevState => {
      return {
        subscribed: (prevState.subscriptions[button].subscribed = false)
      };
    });
  };

  render() {
    //Checks for open connection
    if (socket.readyState === 1) {
      //Connection is open
      socket.onmessage = evt => {
        console.log("RESPONSE: " + evt.data);
        switch(evt) {
          //If evt.data returns error
          case (evt.data.search("error") !== -1 ):
          console.log("evt.data returns error");
          break;

          //if evt.data returns string
          case(evt.data.search("test") !== -1):
          this.setState(prevState => {
              return {
                // TODO: How to handle evt.data?
                feed: [
                  ...prevState.feed,
                  {
                    title: evt.data,
                    time: new Date().toLocaleTimeString()
                  }
                ]
              };
            });
          break;

          //If evt.data returns news json
          case(evt.data.search("news")):
          const newsObj = JSON.parse(evt.data);
          
          /*
            Uncomment below to set feed state. May possible need to format data to feed object. 
            Look above to object state.feed for current reference
          */

          this.setState(prevState => {
            return {
              //Tcase(evt.data.search("news")):ODO: How to handle evt.data?
              feed: [
                ...prevState.feed,
                {
                  title: newsObj.result.abstract,
                  time: new Date().toLocaleTimeString()
                }
              ]
            };
          });
          break;

          //if evt.data returns weather json
          case(evt.data.search("weather")):
          const weatherObj = JSON.parse(evt.data);
          this.setState(prevState => {
            return {
              feed: [
                ...prevState.feed,
                {
                  title: weatherObj.result.abstract,
                  time: new Date().toLocaleTimeString()
                }
              ]
            };
          });
          break;

          default:
          console.log("nothing in response");
          break;
        }
        
      };
    } 
    else {
      //Connection needs to be opened
      socket.onopen = evt => {
        console.log("OPEN");
        // socket.send("subscribe test.channel");
        socket.onmessage = evt => {
          console.log("RESPONSE: " + evt.data);
        };
      };
    }

    console.log("Rendered!");

    return (
      <Grid>
        <PageHeader>
          <Row>
            <Col md={3}>PUB -- SUB</Col>
            <Col md={3}>
              <small>Server by Brandon Manke</small>
            </Col>
            <Col md={3}>
              <small>UI and Sockets by Rob Putyra</small>
            </Col>
          </Row>
        </PageHeader>
        <Row>
          <Col md={8}>
            <Table striped bordered condensed hover responsive>
              <thead>
                <tr>
                  <th className="text-center" colSpan="2">
                    Feed
                  </th>
                </tr>
              </thead>
              <Feed
                feed={this.state.feed}
                subscriptions={this.state.subscriptions}
                // Test with {feed just to see}
              />
            </Table>
          </Col>
          <Col className="text-center" md={4}>
            <Subscriptions
              subscriptions={this.state.subscriptions}
              changeSubscriptions={this.handleSubscription}
              changeUnsubscriptions={this.handleUnsubscription}
            />
          </Col>
        </Row>
      </Grid>
    );
  }
}

export default App;
