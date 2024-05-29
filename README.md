# Land Use Map

This project is in its early stages.

An interactive map of land use in the uk so that we can understand how the land is used, and how it can be changed.

## Setup
Install [Caddy](https://caddyserver.com/) and then `caddy run`.

There are some js scripts to work out whether an h3 cell is in the uk, this requires a google maps key at $MAPS_KEY.

## TODO
- Integrate [land use survey data](https://www.gov.uk/government/statistics/land-use-in-england-2022) to categorise H3 cells.
- Determine suitable H3 cell resolution.
- Display metrics and create interactive element of the model.
