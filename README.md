# LICENSER #

Simple api that retrun text of license with current year and your copyright.

Current licenses:
1. Apache License 2.0 (Apache-2.0)
2. Artistic License 2.0 (Artistic-2.0)
3. 2-clause BSD License (BSD-2-Clause)
3. 3-clause BSD License (BSD-3-Clause)
4. GNU General Public License version 3 (GPL-3.0)
5. ISC License (ISC)
6. GNU Lesser General Public License version 3 (LGPL-3.0)
7. MIT License (MIT)
8. Universal Permissive License (UPL)
9. Do What The Fuck You Want To Public License (WTFPL)

## Build ##
- `git clone git@github.com:k33nice/licenser`
- `cd licenser`
- `go build`

## Deploy ##
- `cp .env.sample .env`
- change enviroment variables
- `make rollout`

## Run ##
`./licenser`

## Usage ##
Make get request with choosen license
1. Apache License 2.0 (Apache-2.0)

    `curl -s http://localhost:33654/apache-2`
2. Artistic License 2.0 (Artistic-2.0)

    `curl -s http://localhost:33654/artistic-2`
3. 2-clause BSD License (BSD-2-Clause)

    `curl -s http://localhost:33654/bsd-2`
3. 3-clause BSD License (BSD-3-Clause)

    `curl -s http://localhost:33654/bsd-3`
4. GNU General Public License version 3 (GPL-3.0)

    `curl -s http://localhost:33654/gpl-3`
5. ISC License (ISC)

    `curl -s http://localhost:33654/isc`
6. GNU Lesser General Public License version 3 (LGPL-3.0)

    `curl -s http://localhost:33654/lgpl-3`
7. MIT License (MIT)

    `curl -s http://localhost:33654/mit`
8. Universal Permissive License (UPL)

    `curl -s http://localhost:33654/upl`
9. Do What The Fuck You Want To Public License (WTFPL)

    `curl -s http://localhost:33654/wtfpl`

### Parameters ###

- `n` -- Name (James Bond)
- `e` -- Email (james.bond@mi6.gov)
- `y` -- Year(s) (2017, 2000 - 2017)
- `p` -- Project (Casino Royale)

### Examples ###

- `curl -s http://localhost:33654/mit`
- `curl -s http://localhost:33654/mit -d e=foo@bar.baz --data-urlencode="2000 - 2017"`
