import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type Mutation = {
  __typename?: 'Mutation';
  addSymbolToWatchList: Symbol;
  removeSymbolFromWatchList: Symbol;
};


export type MutationAddSymbolToWatchListArgs = {
  id: Scalars['ID'];
};


export type MutationRemoveSymbolFromWatchListArgs = {
  id: Scalars['ID'];
};

export type PriceUpdate = {
  __typename?: 'PriceUpdate';
  symbolName: Scalars['String'];
  session: TradingSession;
};

export type Query = {
  __typename?: 'Query';
  watchedSymbols: Array<Symbol>;
  symbol?: Maybe<Symbol>;
};


export type QuerySymbolArgs = {
  id: Scalars['ID'];
};

export type Subscription = {
  __typename?: 'Subscription';
  priceUpdatesFromSymbol: PriceUpdate;
};


export type SubscriptionPriceUpdatesFromSymbolArgs = {
  id: Scalars['ID'];
};

export type Symbol = {
  __typename?: 'Symbol';
  id: Scalars['ID'];
  name: Scalars['String'];
  sessions: Array<TradingSession>;
};


export type TradingSession = {
  __typename?: 'TradingSession';
  time: Scalars['Time'];
  open: Scalars['Int'];
  close: Scalars['Int'];
  high: Scalars['Int'];
  low: Scalars['Int'];
};

export type TradingSessionsFragmentFragment = (
  { __typename?: 'TradingSession' }
  & Pick<TradingSession, 'open' | 'close' | 'high' | 'low' | 'time'>
);

export type SymbolFragmentFragment = (
  { __typename?: 'Symbol' }
  & Pick<Symbol, 'id' | 'name'>
  & { sessions: Array<(
    { __typename?: 'TradingSession' }
    & TradingSessionsFragmentFragment
  )> }
);

export type SymbolByIdQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type SymbolByIdQuery = (
  { __typename?: 'Query' }
  & { symbol?: Maybe<(
    { __typename?: 'Symbol' }
    & SymbolFragmentFragment
  )> }
);

export const TradingSessionsFragmentFragmentDoc = gql`
    fragment TradingSessionsFragment on TradingSession {
  open
  close
  high
  low
  time
}
    `;
export const SymbolFragmentFragmentDoc = gql`
    fragment SymbolFragment on Symbol {
  id
  name
  sessions {
    ...TradingSessionsFragment
  }
}
    ${TradingSessionsFragmentFragmentDoc}`;
export const SymbolByIdDocument = gql`
    query SymbolByID($id: ID!) {
  symbol(id: $id) {
    ...SymbolFragment
  }
}
    ${SymbolFragmentFragmentDoc}`;

/**
 * __useSymbolByIdQuery__
 *
 * To run a query within a React component, call `useSymbolByIdQuery` and pass it any options that fit your needs.
 * When your component renders, `useSymbolByIdQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSymbolByIdQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useSymbolByIdQuery(baseOptions: Apollo.QueryHookOptions<SymbolByIdQuery, SymbolByIdQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<SymbolByIdQuery, SymbolByIdQueryVariables>(SymbolByIdDocument, options);
      }
export function useSymbolByIdLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<SymbolByIdQuery, SymbolByIdQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<SymbolByIdQuery, SymbolByIdQueryVariables>(SymbolByIdDocument, options);
        }
export type SymbolByIdQueryHookResult = ReturnType<typeof useSymbolByIdQuery>;
export type SymbolByIdLazyQueryHookResult = ReturnType<typeof useSymbolByIdLazyQuery>;
export type SymbolByIdQueryResult = Apollo.QueryResult<SymbolByIdQuery, SymbolByIdQueryVariables>;