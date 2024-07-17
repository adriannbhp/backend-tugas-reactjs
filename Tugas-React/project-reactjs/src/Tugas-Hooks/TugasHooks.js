import React, { useState, useEffect } from 'react';

class CurrentTime extends React.Component {
    constructor(props) {
        super(props);
        this.state = { time: new Date() };
    }

    componentDidMount() {
        this.interval = setInterval(() => this.setState({ time: new Date() }), 1000);
    }

    componentWillUnmount() {
        clearInterval(this.interval);
    }

    render() {
        return <div className="time-display">Now At - {this.state.time.toLocaleTimeString()}</div>;
    }
}

const Countdown = ({ start, onComplete }) => {
    const [count, setCount] = useState(start);

    useEffect(() => {
        if (count > 0) {
            const timer = setInterval(() => {
                setCount(count => count - 1);
            }, 1000);
            return () => clearInterval(timer);
        } else {
            onComplete();
        }
    }, [count, onComplete]);

    return <div className="countdown-display">Countdown: {count}</div>;
};

const TugasHooks = () => {
    const startCount = 100;
    const [showComponents, setShowComponents] = useState(true);

    const handleComplete = () => {
        setShowComponents(false);
    };

    return (
        <div>
            {showComponents && (
                <>
                    <CurrentTime />
                    <Countdown start={startCount} onComplete={handleComplete} />
                </>
            )}
        </div>
    );
};

export default TugasHooks;
