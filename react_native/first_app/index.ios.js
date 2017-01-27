// Import library for component
import React from 'react';
import { AppRegistry } from 'react-native';

// Import our own Components
import Header from './src/components/header';

// Create a Component
const App = () => (
  <Header headerText={'Albums!'}/>
);

// Render it to the device
AppRegistry.registerComponent('first_app', () => App);
