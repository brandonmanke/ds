import React, { Component } from "react";
import "./App.css";
import "./index.css";

import { Table, Grid, Row, Col, PageHeader } from "react-bootstrap";
import Feed from "./components/Feed";
import Subscriptions from "./components/Subscriptions";

//Establishing socket instance
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onclose = evt => {
  console.log("CLOSE");
};

socket.onerror = evt => {
  console.log("ERROR: " + evt.data);
};

//Variables for dummy feed data
var date = new Date();
var time = date.getTime();

var hour = 60 * 60 * 1000;

const times = [
  time,
  time + 3 * hour,
  time + 6 * hour,
  time + 1 * hour,
  time + 4 * hour,
  time + 7 * hour,
  time + 2 * hour,
  time + 5 * hour
];

class App extends Component {
  state = {
    // Dummy feed data. Will switch to null once ready to handle socket messages
    feed: [
      {
        type: "weather",
        title: "Weather1",
        time: times[0]
      },
      {
        type: "weather",
        title: "Weather2",
        time: times[1]
      },
      {
        type: "weather",
        title: "Weather3",
        time: times[2]
      },
      {
        type: "news",
        title: "News1",
        time: times[3]
      },
      {
        type: "news",
        title: "News2",
        time: times[4]
      },
      {
        type: "news",
        title: "News3",
        time: times[5]
      },
      {
        type: "friends",
        title: "Friend1",
        time: times[6]
      },
      {
        type: "friends",
        title: "Friend2",
        time: times[7]
      }
    ],
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

  // Might need to use this. Placeholder for now
  componentDidMount() {}

  //Handles button controls to subscribe to channel
  handleSubscription = button => {
    console.log("frontend subscribed")
    socket.send("subscribe test.channel");
    this.setState(prevState => {
      return {
        prevState: (prevState.subscriptions[button].subscribed = true)
      };
    });
  };

  //Handles button for unsubscribe
  handleUnsubscription = button => {
    console.log("frontend unsubscribed")
    socket.send("unsubscribe test.channel");
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
          //If evt.data returns data
          default:
          /*
            Uncomment below to set feed state. May possible need to format data to feed object. 
            Look above to object state.feed for current reference
          */

          // this.setState(prevState => {
          //   return {
          //     //TODO: How to handle evt.data?
          //     // subscribed: (prevState.feed = evt.data)
          //   };
          // });
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
