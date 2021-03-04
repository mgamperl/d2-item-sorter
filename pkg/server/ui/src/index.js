import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter } from 'react-router-dom';
import {Grommet} from "grommet";
import {theme} from "./theme.js"

import { createOvermind} from 'overmind'
import { Provider } from 'overmind-react'
import { config } from './overmind'

const overmind = createOvermind(config)

ReactDOM.render(
  <React.StrictMode>
     <Provider value={overmind}>
    <Grommet themeMode="dark" theme={theme} >
      <BrowserRouter >
        <App />
      </BrowserRouter>
    </Grommet>
    </Provider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
