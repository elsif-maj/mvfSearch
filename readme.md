# Readme

This app is currently aimed at being integrated with Umbra ( https://github.com/elsif-maj/Umbra-VPS-Deployment ), but I will eventually generalize the APIs and offer configurability for it to be as broadly applicable as I can make it.

At present the app will take primary keys of documents containing text in a database, ingest the text content, and route it through a tokenization and ngram-ification process. The resultant string arrays can then be processed in a way that they are fed into Redis (although I am looking to make this interface generalized) to create a reverse index which maps keys that include user IDs, concatenated with the tokens and ngrams to the document IDs (or primary keys) from which they were generated, leading to one large key-value store with all tokens and token ngrams (from length 2 to 5 currently) from all documents in the database being look-up-able in 0(1) time in a way that is indexed to the users who are associated with the documents. If one pipes all relevant documents through this process the result is individualized full text search (yielding document ID), indexed per user, for all words and sequences of words (from 2 to 5) that show up in their documents.

A list of stuff that I'd like to keep an awareness of as I continue to work:


-Fix error handling

-Add stop word removal in the indexing flows

-API Keys

-Standardization of naming for functions/parameters/variables

-Set up env/secrets for Redis connection

-Tests

-Docs

-Consider what API design should be for a more app-neutral service (i.e. generalizing away from plugging this in to Umbra)

-Consolidate packages 

-Prefix search?





