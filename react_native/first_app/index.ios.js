// Import library for component
import React from 'react';
import { AppRegistry, Text } from 'react-native';
// Create a Component

const App = () => (
  <Text>Some Text</Text>
);

// Render it to the device
AppRegistry.registerComponent('first_app', () => App);
