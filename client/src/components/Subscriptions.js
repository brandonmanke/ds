import React from 'react';

import { Row, Button } from 'react-bootstrap';

class Subscriptions extends React.Component {

    
    render() {

        const {
            subscriptions,
            changeSubscriptions
        } = this.props;

        return(
            <div className = "subscriptions">
          <Row className="with-margin">
            <Button 
                block
                active = {subscriptions.default}
                onClick={() => changeSubscriptions("default")}>
                Default
            </Button>
          </Row>
          <Row className="with-margin">
            <Button 
                block bsStyle="primary"
                active = {subscriptions.weather}
                onClick={() => changeSubscriptions("weather")}>
                Weather
            </Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates a successful or positive action */}
            <Button block bsStyle="success"
                active = {subscriptions.news}
                onClick={() => changeSubscriptions("news")}>
                News
            </Button>
          </Row>
          <Row className="with-margin">
            {/* Contextual button for informational alert messages */}
            <Button block bsStyle="info"
                active = {subscriptions.friends}
                onClick={() => changeSubscriptions("friends")}>
                Friends
            </Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates caution should be taken with this action */}
            <Button block bsStyle="warning">Warning</Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates a dangerous or potentially negative action */}
            <Button block bsStyle="danger">Danger</Button>
          </Row>
          </div>
        )
    }
}

export default Subscriptions;