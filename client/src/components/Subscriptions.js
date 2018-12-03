import React from 'react';

import { Row, Button } from 'react-bootstrap';


class Subscriptions extends React.Component {

    
    render() {

        const {
            subscriptions,
            changeSubscriptions,
            changeUnsubscriptions
        } = this.props;

        return(
        <div className = "subscriptions">
          <Row className="with-margin">
            <Button 
                block bsStyle="danger"
                active = {subscriptions[0].subscribed}
                onClick={subscriptions[0].subscribed ? 
                    () => changeUnsubscriptions(0) : () => changeSubscriptions(0) }>
                Test
            </Button>
          </Row>
          <Row className="with-margin">
            <Button 
                block bsStyle="primary"
                active = {subscriptions[1].subscribed}
                onClick={subscriptions[1].subscribed ? 
                    () => changeUnsubscriptions(1) : () => changeSubscriptions(1) }>
                Weather
            </Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates a successful or positive action */}
            <Button block bsStyle="success"
                active = {subscriptions[2].subscribed}
                onClick={subscriptions[2].subscribed ? 
                    () => changeUnsubscriptions(2) : () => changeSubscriptions(2) }>
                News
            </Button>
          </Row>
        </div>
        )
    }
}

export default Subscriptions;