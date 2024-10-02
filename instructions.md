# [G]ithub[X]plore TUI/CLI (gx)

```
TIMEBOX:    4-6 hours max. We mean it! Set at timer and hard-stop at 6 hours ⏱
LANGUAGES:  Go, Python, Typescript, Rust, C/C++
TESTS:      nice to have, but not mandatory (at least one unit test)
DOCS:       nice to have, but not mandatory
```

## Description
`gx` TUI/CLI is a tool to discover and explore GitHub repos. Think of it as  CLI tool that cooresponds to the "Explore" section on github.com. You can give it tags/topics, and it will return back results that match the search criteria. 

Results can be retuned as a single static list of `N` many items, or as an "Text/Terminal User Interface" interactive list.

Your task is to build the `gx` TUI/CLI.

## Design
You are free to use whatever TUI/CLI design elements you wish.

### TUI/CLI
Terminal User Interfaces are a unique kind of CLI application which uses interfactive text based elements to create a UI on the command line. This includes text boxes, buttons, scrollable lists, etc.

### Repo properties
The returned items from a given search should have the following repo properties:
 - Name
 - Owner
 - Primary language
 - Description
 - Number of Stars
 - Number of Watchers
 - Number of issues
 - Numer of PRs
 - Number of total contributors
 - List of tags

### Related repos
The TUI interactive list of repos from the search results should also allow one to explore "related" repos of the target one. This should take into account owner, tags, and language to produce another list of repos, which can also be explored through. Similar to a knowledge graph jumping between related topics, but with GitHub repos.

You are free to implement any other feature you wish, however these are the core required ones. If you are unable to finish these core features, that is not a problem, as we are more concerned about the structure of your code, then the number of features implemented. See the [Assessment][#Assessment].

## Dependency
`gx` will need to use the Github GraphQL API extensively, so it will be necessary to learn and understand the documentation. This can be found at [https://docs.github.com/en/graphql](https://docs.github.com/en/graphql)

## Example Output

```
> gx --tag kubernetes

|======================================================================|
| Search: tag:kubernetes                                               |
|======================================================================|
|    NAME    |    OWNER   |  STARS  |  # of Issues  |  # of PRs  | ... |
|======================================================================|
|            |            |         |               |            |     |
| Kubernetes | Kubernetes |  84.1K  |      1.7K     |    806     | ... |
|            |            |         |               |            |     |
|-----------------------------------------------------------------------
|            |            |         |               |            |     |
| > netdata  |   netdata  |  57.2K  |      162      |    104     | ... |
|            |            |         |               |            |     |
|-----------------------------------------------------------------------
|            |            |         |               |            |     |
|    etcd    |  etcd-io   |  38.4K  |      130      |     60     | ... |
|            |            |         |               |            |     |
|-----------------------------------------------------------------------
|            |            |         |               |            |     |
| hashicorp  |   consul   |  23.8K  |      906      |    153     | ... |
|            |            |         |               |            |     |
|-----------------------------------------------------------------------
|                                                                      |
|  ...                                                                 |
|                                                                      |
|======================================================================|
|                                                             | EXIT-> |
|======================================================================|
```

Using TUI based UI/UX elements, you should be able to *select* a repo to further explore through.

> Note: the `>` next to `netdata` is a interactive TUI selector, to select
> the next item to explore *into*.

Once the `netdata` repo row is selected, we return the next group of repos that match that criteria

```
|========================================================================|
| Search: related:netdata/netdata                                        |
|========================================================================|
|    NAME     |    OWNER    |  STARS  |  # of Issues  |  # of PRs  | ... |
|========================================================================|
|             |             |         |               |            |     |
|    huginn   |   huginn    |  32.9K  |      493      |    82      | ... |
|             |             |         |               |            |     |
|-----------------------------------------------------------------------
|             |             |         |               |            |     |
| sweetalert2 | sweetalert2 |  14.0K  |      6        |      1     | ... |
|             |             |         |               |            |     |
|-----------------------------------------------------------------------
|             |             |         |               |            |     |
|  containrrr | watchtower  |   9.6K  |      53       |     7      | ... |
|             |             |         |               |            |     |
|------------------------------------------------------------------------|
|                                                                        |
|  ...                                                                   |
|                                                                        |
|========================================================================|
| <-BACK |                                                      | EXIT-> |
|========================================================================|
```

## Requirements
- Submitted code
- At least one unit test
- `README.md` file containing:
    - A short explanation of what you built
    - How to test/demo/run (if applicable)
        - NOTE: a 'working' example/client is awesome, however it is NOT a hard requirement. We mean it!
    - Any feedback/notes (i.e. if something was hard, confusing, frustrating, etc)
    - Anything else you'd lke us to know about your submission
- `ROADMAP.md` file  with what you would add/change if you had more time. Dream big.

## Assessment
Your code will be assessed using the same goals and requirements that we use day to day at Source. On its overall design, structure, and readability. **We rather you implement less features, that are well thought out, then many just for the sake of completeness.**

### Readability 
Readability is an important factor, as we try to follow the concepts and goals of *Literate Programming*. 

> *Literate Programming*:
>
> Instead of imagining that our main task is to instruct a computer what to do, let us concentrate rather on explaining to human beings what we want a computer to do.

This is a subjective goal to evaluate, but something we look, as we believe it is important to effectively communicate to your peers and collaborators what the goal of a specific program/function/snippet is. This is even more important in remote/hybrid work forces, where asynchronos work is common, and you don't always have the freedom to casually talk to your coworker beside you.

### Abstractions and Structure
The common buzzword philisophy for Starups is to "Move fast, and break things". In certain environments, this can be an effective stratedgy. However at Source, we take a more measured and slower approach to most things we do.

Our goals and engineering challenges here at Source often take us to the bleeding edge of various technologies, and when walking on new ground, we want to be confident that the steps we are taking are the rights ones, so we may have a better foundation to build on later. 

It's one thing if your ToDo SaaS app goes offline for an hour, but a Database that irevocably destroys or losses data is a non-starter

As such, running code on a machine is usually the last step in the engineering process here at Source, which starts with ideation and design based on clearly defined goals.

## Need Help?
If you need any help, want to ask a question, need clarification, or anything else, feel free to join our Community Discord to chat with our developers. Or, simply open up an issue thread on this repo. Both are suggested.

If you are stuck, and have reached out for help either via [Issues](https://github.com/sourcenetwork/hiring/issues) or [Discord](https://discord.gg/57UNewmXtE), feel free to put the project away untill you have got your answer. As there is a working time limit of 6 hours, and we don't want you to waste it running in circles on a problem or clarification. 

> Hint: Some of this document is intentionally vague

## Background
> Note: You are free to implement the TUI/CLI design however you wish, and the above examples should be used as inspiration.

### Terminal User Interface
 - https://en.wikipedia.org/wiki/Text-based_user_interface
 - https://github.com/rothgar/awesome-tuis
 - https://appliedgo.net/tui/