import * as React from 'react';
import './App.css';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Receipt from './components/Receipt/Receipt'

const logo = require('./logo.svg');

class App extends React.Component<{}, {}> {
  render() {
    return (
        <MuiThemeProvider>
            <div className="App">
                <div className="App-header">
                    <img src={logo} className="App-logo" alt="logo" />
                    <h2>Welcome to Mars</h2>
                </div>
                <div>
                    <Receipt/>
                </div>
            </div>
        </MuiThemeProvider>
    );
  }
}

export default App;
