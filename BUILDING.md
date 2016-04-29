# Building Arukas CLI

This document contains details about the process for building binaries for Arukas CLI

## QuickBuild

**Please note: Replaced by your arukas token and aruaks api secret is
 `YOUR_API_TOKEN` and `YOUR_API_SECRET`**

* Clone the repo: `git clone https://github.com/arukasio/cli.git`
* CLI Build: `docker build -t arukasio/arukas:patch .`
* Test execute the CLI: `docker run --rm -e ARUKAS_JSON_API_TOKEN="YOUR_API_TOKEN"
-e ARUKAS_JSON_API_SECRET="YOUR_API_SECRET" arukasio/arukas:patch`
