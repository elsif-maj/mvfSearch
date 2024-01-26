# Readme

Alright, alright. I really need to start documenting things -- really! I'm filled with admiration for the lovely docs that I've been seeing now that I have more time to browse through projects on GitHub. I'd like to follow suit. 2024 is the year of docs. And tests.

But that day is not today. I'm still hacking away on this exploratory project, and there is much to do before committing any words to a page about what the *state of things* is. I can, however, now grab docId search term matches at the word level and higher -- ngrams of words as well. O(1) btw -- nbd.

There is still a fair bit of Go noobishness to be cleaned up. Error handling and package proliferation are setting off alarm bells for me -- and that's just stuff I know enough to recognize. The API is untouched since day 1, and there was no plan for it on day 1.

All that being said, I'm going to at least start keeping in this readme a list of stuff that I'd like to keep an awareness of as I continue to work:

-API Keys

-Standardization of naming for functions/parameters/variables

-Add stop word removal in the indexing flows

-Consider what API design should be for a more app-neutral service (i.e. generalizing away from plugging this in to Umbra)

-Prefix search?

-Set up env/secrets for Redis connection
