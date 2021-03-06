// todo: add hits limit depending on the api subscription plan (60 for free)
// use n for loading too, to limit the number of cards loaded from the api
// handle 403 errors in case we are making too many requests somehow
// change load to "prefetch", change load flag so that it simply adds words
// to the local db as pending words
// change export feature so that it if there are no ready to export cards it will try to
// prefetch pending words
// effectively, you would rarely want to run "prefetch" on its own, and prefer adding words
// in pending state and let export to run "prefetch" for you
// but, in case you want to prefetch your words

// add import command, accept newline separated list of words from stdin
// implement import from kindle db
// add blacklist command that disallows a word to be exported, and never tries to fetch it
// add reset-export command that takes a list of words and marks them as not exported
// add enqueue command that takes a list of words and adds them to the next export queue
// add export command that exports to anki deck. Take words from queue first, then from the
// regular word collection, provided word is not marked as already exported. Mark all successfully
// exported words
// mark words that were failed to fetch (404) to avoid refetching them the next export

// implement kindle db import
