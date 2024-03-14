import {
  ApolloClient,
  //InMemoryCache /*, createHttpLink */,
} from '@apollo/client/core';
// import { ApolloClients } from '@vue/apollo-composable';
import { boot } from 'quasar/wrappers';
import { getClientOptions } from '../apollo';
import { provideApolloClient } from '@vue/apollo-composable';


export default boot(
  /* async */ ({ /* app */ }) => {
    //const cache = new InMemoryCache();

    // Default client.
    const options = /* await */ getClientOptions();
    const apolloClient = new ApolloClient(options);

    // // Additional client `clientA`
    // const optionsA = { ...options }
    // // Modify options as needed.
    // optionsA.link = createHttpLink({ uri: 'http://clientA.example.com' })
    // const clientA = new ApolloClient(optionsA)

    // // Additional client `clientB`
    // const optionsB = { ...options }
    // // Modify options as needed.
    // optionsB.link = createHttpLink({ uri: 'http://clientB.example.com' })
    // const clientB = new ApolloClient(optionsB)

    /* const apolloClients: Record<string, ApolloClient<unknown>> = {
      default: apolloClient,
      // clientA,
      // clientB,
    }; */

    provideApolloClient(apolloClient);
    // provideApolloClients(apolloClient,{default: apolloClient,clientA: apolloClientA});

    // app.provide(ApolloClients, apolloClients);
    // app.provide(ApolloClients, {default: apolloClient});
  }
);
