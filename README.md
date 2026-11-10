# Pokedex

A terminal-based Pokédex and exploration game written in Go. Browse the Pokémon world's location areas, discover which Pokémon live there, try to catch them, and inspect the ones you've collected — all from an interactive REPL powered by [PokéAPI](https://pokeapi.co/).

```
Pokedex > map
Location Areas:
  - canalave-city-area
  - eterna-city-area
  - oreburgh-city-area
  ...

Pokedex > explore eterna-city-area
Exploring eterna-city-area:
  Found Pokemon:
    - gastly
    - abra
    - kadabra
    ...

Pokedex > catch abra
Throwing a Pokeball at abra...
abra was caught!

Pokedex > inspect abra
Name: abra
Height: 9
Weight: 195
Stats:
  - hp: 25
  - attack: 20
  - defense: 15
  ...

Pokedex > pokedex
Your Pokedex:
  - abra
```

## Features

- **Interactive REPL** with a built-in `help` command and friendly prompts.
- **Location exploration** — page forward and backward through location areas, then list the Pokémon found in each one.
- **Catching** — attempt to catch any Pokémon; tougher Pokémon (higher base experience) are harder to catch.
- **Personal Pokédex** — caught Pokémon are stored in memory and can be listed and inspected (name, height, weight, and base stats).
- **TTL response cache** — API responses are cached in memory for an hour, so repeated lookups are instant and don't hit the network.
- **Graceful HTTP handling** — request timeouts, non-200 status checks, and JSON decoding errors are reported without crashing the REPL.

## Commands

| Command | Description |
| --- | --- |
| `help` | Show the list of available commands |
| `map` | List the next page of location areas |
| `mapb` | List the previous page of location areas |
| `explore <location_area>` | List the Pokémon found in a location area |
| `catch <pokemon_name>` | Attempt to catch a Pokémon |
| `inspect <pokemon_name>` | Show details of a Pokémon you've caught |
| `pokedex` | List the names of all caught Pokémon |
| `exit` | Exit the Pokédex |

Input is case-insensitive, and extra whitespace is ignored, so `CATCH Pikachu` works the same as `catch pikachu`.

## Requirements

- Go **1.27.3** or newer (see `go.mod`)
- An internet connection (all data comes from `https://pokeapi.co/api/v2`)

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/abdullahsamir74/pokedex.git
cd pokedex
go build -o pokedex .
./pokedex
```

Or install it directly with Go:

```bash
go install github.com/abdullahsamir74/pokedex@latest
pokedex
```

## Usage

Start the program and type commands at the `Pokedex >` prompt. A typical session:

1. `map` to see the first page of location areas.
2. `map` / `mapb` to page through the list.
3. `explore <area>` to see which Pokémon appear there. Names are lowercase and hyphenated, e.g. `explore eterna-city-area`.
4. `catch <pokemon>` to try to catch one.
5. `pokedex` to see what you've caught, and `inspect <pokemon>` for its details.
6. `exit` (or `Ctrl-D`) to quit.

### Catching mechanics

A catch attempt rolls a random number between `0` and the Pokémon's base experience. If the result is above `50`, the Pokémon escapes; otherwise it's caught. In practice this gives roughly a `50 / base_experience` chance (capped at 100%), so common Pokémon are easy and legendaries are rare.

### Cache behavior

Each unique API URL is cached for one hour. When a cached response is used, a `Cache hit for <url>` line is printed before the result. This applies to location area lists, individual location areas, and Pokémon lookups.

## Project structure

```
.
├── main.go                      # Entry point: wires the client and starts the REPL
└── internal/
    ├── repl/                    # Interactive shell and game commands
    │   ├── repl.go              # Loop, command registry, input parsing
    │   ├── config.go            # Shared state (client, pagination, caught Pokémon)
    │   ├── commands.go          # All command callbacks
    │   └── repl_test.go
    ├── pokeapi/                 # PokéAPI client and response types
    │   ├── client.go            # HTTP client with caching
    │   ├── location_area.go     # Location area endpoints and types
    │   ├── pokemon.go           # Pokémon endpoint and type
    │   └── client_test.go
    └── pokecache/               # Thread-safe TTL cache with background reaper
        ├── pokecache.go
        └── pokecache_test.go
```

### How it works

- **REPL loop** (`internal/repl/repl.go`) reads a line, normalizes it with `cleanInput` (lowercase + whitespace split), looks up the first token in a command registry, and calls the matching callback with any remaining arguments. `help` lists commands alphabetically, and an unknown command prints a hint instead of exiting.
- **Shared state** lives in `repl.Config`: the API client, the next/previous pagination URLs, and the in-memory collection of caught Pokémon.
- **API layer** (`internal/pokeapi`) wraps `net/http` with a one-minute timeout. A single `get` helper checks the cache, falls back to the network, validates the status code, decodes JSON, and stores successful responses as raw bytes. The three endpoint methods (`ListLocationAreas`, `GetLocationArea`, `GetPokemon`) are thin wrappers over it, and names are URL-escaped before being sent.
- **Cache layer** (`internal/pokecache`) is a mutex-protected `map` of byte slices with creation timestamps. A background goroutine removes entries older than the configured interval.

## Testing

```bash
go test ./...
```

Tests cover input parsing (`cleanInput`), the command registry, the HTTP client (caching, non-200 responses, invalid JSON) using `httptest`, and cache behavior (add/get and TTL reaping).

To check formatting and static analysis:

```bash
gofmt -l .
go vet ./...
```

## Notes and limitations

- Caught Pokémon are stored **in memory only** — they are lost when the program exits. There is no save file or persistence yet.
- No authentication or API key is required; PokéAPI is free and public.
- Pokémon and location area names must match PokéAPI's slugs (lowercase, hyphenated), e.g. `mr-mime` rather than `Mr. Mime`.

## Acknowledgements

- Pokémon data from [PokéAPI](https://pokeapi.co/).
- Pokémon and Pokémon character names are trademarks of Nintendo / Game Freak / Creatures Inc. This project is a non-commercial learning exercise.

## License

Released under the [MIT License](LICENSE).

