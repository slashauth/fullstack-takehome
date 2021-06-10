import React from 'react';
import ReactDOM from 'react-dom';
import {
  ApolloProvider,
  HttpLink,
  ApolloClient,
  InMemoryCache,
} from "@apollo/client";
import './index.css';
import App from './App';

const graphQLink = new HttpLink({
  uri: "http://localhost:8080/query",
});

const cache = new InMemoryCache();
const client = new ApolloClient({
  link: graphQLink,
  cache,
  credentials: "include",
  resolvers: {},
});

ReactDOM.render(
  <ApolloProvider client={client}>
      <App />
  </ApolloProvider>,
  document.getElementById('root'),
);
