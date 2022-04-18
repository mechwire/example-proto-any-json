# Example: Proto-Any-JSON

As a disclaimer, I didn't think this was the right way to use the technologies. (I don't think this solution was great either.) But, it fit the requirements, was a good learning experience, and took a lot of time.

--

This is a repository that holds the cumulative work from the following:
* [example-proto-types](https://github.com/jncmaguire/example-proto-types), which holds proto definitions
* [example-proto-custom-resolver](https://github.com/jncmaguire/example-proto-custom-resolver), which holds proto definitions and a resolver to recognize them

Overall, this is part of exploratory code given the following niche requirements:

* JSON objects must originate from proto definitions
* JSON objects should be deserialized into a map, where the map has typed objects. (The default map behavior would be nested `map[string]interface{}`)
* The place in which they will be used in the code will take various types of proto types.
* The consumer and the multiple producers of the proto types are maintained by different groups and intentionally decoupled.

As such, this solution provides the following:
* Semi-independent yet centralized message management with Go submodules / versioning
* A centralized resolver

This repo shows, overall, how to take the JSON representation of an Any object and create a map from it.