# Language Server Protocol

A Language Server is meant to provide the language-specific smarts and communicate
with development tools over a protocol that enables inter-process communication.

The idea behind the Language Server Protocol (LSP) is to standardize the protocol
for how such servers and development tools communicate.
This way, a single Language Server can be re-used in multiple development tools,
which in turn can support multiple languages with minimal effort.

[More information about it...](https://microsoft.github.io/language-server-protocol/)

## Index

- [Setup](#setup)
- [What is this practice about?](#what-is-this-practice-about%3F)

### Setup

Create a fork of this repository and clone it locally

```bash
git clone https://github.com/<username>/lsp-practice # or us git ssh
# or with gh-cli
gh repo clone <username>/lsp-practice
```

Install dependencies

```bash
go install mvdan.cc/gofumpt@latest
```

```bash
npm install # or pnpm/yarn/bun ...
npm build # or pnpm/yarn/bun ...
# manually
go build main.go
```

Make a connection to your client, in my case NeoVim:

```bash
cd <neovim-config>
touch after/plugin/load_lsp_practice.lua
```

```lua
local client = vim.lsp.start_client({
    name = "lsp_practice",
    cmd = "path/to/your/lsp-practice/main",
    on_attach = function(client, bufnr)
        -- your on_attach behaviour, mappings, client capabilities ...
    end,
})

if not client then
    -- vim.notify to show any error how
    return
end

-- attachs client to markdown files
vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

### What is this practice about?

Implementing and getting to know the ins and outs of this protocol, to get an lsp
for [HTMX](https://htmx.org/) in as many editors as it can be.

For this implementation, I will be using **Go**.
