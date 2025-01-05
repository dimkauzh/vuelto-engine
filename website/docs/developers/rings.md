# 💍 Vuelto's rings approach

Hey and welcome to this small explanation how Vuelto's 3 ring Engine approach works, and how you can help and improve it! Lets start off with a very easy explanation what this is; this is an implementation how Vuelto goes from a system library to a full fledges API thats being used in the Engine. We call it the "3 Ring Engine approach".

## 🔘 3 Rings

Vuelto is build up from 3 rings. An important concept in this structure is that a ring above is always dependant on the ring below. This is very important, because if ring 1 is not functioning correctly, then ring 2 and 3 won't function correctly either.

We refer to Ring 1 and Ring 2 as "The Underground", while Ring 3 as "The Engine", or "The top layer".

The ring are build in this structure:

- Ring 1: Bare metal API, interacts with the system libraries directly (or Go bindings of them). This is a basic wrapper around the bare metal system libraries.
- Ring 2: The multipurpose API: The API that Vuelto itself is gonna use. This API mixes all the Bare Metal API's in a way, that it comes out as a multi-platform and powerful API.
- Ring 3: The Engine itself: The top layer, the last layer. This is the API that the user is going to interact with. This API must follow the semantic version to not cause chaos.

## 🔢 Versioning

This is a very important part of the rings approach. Vuelto, as every project, can introduce breaking changes. This can happen in both the underground and the Engine. Only these two are really different in the breaking changes approach:

- The Underground: The underground can introduce breaking changes, without worries. But of course, The Engine needs to be ported to support these changes.
- The Engine: In here, things get more strictly. Breaking changes may only be introduces under MAJOR[^1] release. Deprecations are preferably made under MAJOR[^1] releases, but MINOR[^1] is also allowed.

[^1]: Vuelto uses [semantic versioning](https://semver.org/), as MAJOR.MINOR.PATCH.
