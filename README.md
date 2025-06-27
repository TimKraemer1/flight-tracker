# flight-tracker

Welcome aboard **flight-tracker** - a miniature *terminal-based* air traffic control center built with Go! This project lets you fetch information on airports and track historical arrival and departure flights from specified airports. Utilizing the Open-Sky API, you can get the most up-to-date historical flight information for departures and arrivals, fetched and refreshed daily!

___ 

## Overview

**flight-tracker** is a Go-based application designed to monitor, track, and report on flights. It's built to be fast, reliable, and fun to use with a retro-terminal UI design to go with it. Utilizing internal caching with MySQL, API rate limits are never reached for consecutive requests.