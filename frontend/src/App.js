import React, {Component} from 'react';
import Ideas from './components/ideas';


class App extends Component {
  state = {
    ideas: []
  }
  componentDidMount() {
    fetch('http://localhost:9090/')
    .then(res => res.json())
    .then((data) => {
      this.setState({ ideas: data })
    })
    .catch(console.log)
  }
  render () {
    return (
      <Ideas ideas={this.state.ideas} />
    );
  }
}

export default App;
