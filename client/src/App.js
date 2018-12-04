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
        //console.log("RESPONSE: " + evt.data);
        console.log('event:', evt.data)
        //switch(evt.data) {
          //If evt.data returns error

          if (evt.data.includes('error:')) {
            console.log("evt.data returns error");
          } else if (evt.data.includes('test.channel:')) {
            this.setState(prevState => {
              return {
                // TODO: How to handle evt.data?
                feed: [
                  {
                    title: evt.data,
                    time: new Date().toLocaleTimeString()
                  },
                  ...prevState.feed
                ]
              };
            });
          } else if (evt.data.includes('news:')) {
            const str = evt.data
            const newsObj = JSON.parse(str.substr(str.indexOf(':') + 1));
            const rand = Math.floor(Math.random() * newsObj.num_results)
          
            /*
              Uncomment below to set feed state. May possible need to format data to feed object. 
              Look above to object state.feed for current reference
            */

            this.setState(prevState => {
              return {
                //Tcase(evt.data.search("news")):ODO: How to handle evt.data?
                feed: [
                  {
                    title: `Title: ${newsObj.results[rand].title}\n Summary: ${newsObj.results[rand].abstract}`,
                    time: new Date().toLocaleTimeString()
                  },
                  ...prevState.feed
                ]
              };
            });
          } else if (evt.data.includes('weather:')) {
            const str = evt.data
            // Message<test.channel: every 20 seconds>
            const weatherObj = JSON.parse(str.substr(str.indexOf(':') + 1));
            this.setState(prevState => {
              return {
                feed: [
                  {
                    title: `Temperature: ${weatherObj.currently.apparentTemperature}. 
                            Summary: ${weatherObj.currently.summary}. 
                            Lat: ${weatherObj.latitude}, Long: ${weatherObj.longitude}.
                            Timezone: ${weatherObj.timezone}.`,
                    time: new Date().toLocaleTimeString()
                  },
                  ...prevState.feed
                ]
              };
            });
          } else {
            console.log("nothing in response");
          }
        //}
        
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
