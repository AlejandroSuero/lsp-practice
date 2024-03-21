# Language Server Protocol

A Language Server is meant to provide the language-specific smarts and communicate
with development tools over a protocol that enables inter-process communication.

The idea behind the Language Server Protocol (LSP) is to standardize the protocol
for how such servers and development tools communicate.
This way, a single Language Server can be re-used in multiple development tools,
which in turn can support multiple languages with minimal effort.

[More information about it...](https://microsoft.github.io/language-server-protocol/)

## What is this practice about?

Implementing and getting to know the ins and outs of this protocol, to get an lsp
for [HTMX](https://htmx.org/) in as many editors as it can be.

For this implementation, I will be using **Go**.
