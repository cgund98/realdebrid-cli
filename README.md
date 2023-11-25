# Real Debrid CLI

This tool is a simple command-line interface for Real Debrid. You can use it to manage and download files instead of the website.

## Installation

Install the latest binary from the releases page of this repository.

## Examples

### Download from a restricted link
```bash
# set file to download
export REAL_DEBRID_API_TOKEN="myapitoken"
export LINK="https://link.to/my/file.zip"

# Download a file or folder
realdebrid downloads \
    $LINK \
    -o $HOME/Downloads/debrid
```