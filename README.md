# Build a Blog Aggregator in go

This is a CLI tool that allows users to

* Add RSS feeds from across the internet to be collected
* Store the collected posts in a PostgreSQL database
* Follow and unfollow RSS feeds that other users have added
* View summaries of the aggregated posts in the terminal, with a link to the full post

On windows, the program was only tested using WSL.
You will need to install postgresql , golang and goose in order to run the app.

## install postgresql

### linux/wsl (debian)

Install from the package manager
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

Update the password
```
sudo passwd postgres
```

Start the postgres server in the background
```
sudo service postgresql start
```

Enter the psql shell
```
sudo -u psql postgres
```

You should see a new prompt that looks like this:

```
postgres=#
```

Create a new database
```
CREATE DATABASE gator;
```

Connect to the new database
```
\c gator
```

you should see a new prompt that looks like this:
```
gator+#
```

Set the user password (linux/wsl only)
```
ALTER USER postgres PASSWORD 'pickapassword';
```

you can query the database to test
```
SELECT version();
```

you can type `exit` to leave psql shell

### mac os

install using homebrew
```
brew install postgresql@15
```

start the postgres server in the background
```
brew services start postgresql@15
```

Enter the psql shell
```
psql postgres
```

You should see a new prompt that looks like this:

```
postgres=#
```

Create a new database
```
CREATE DATABASE gator;
```

Connect to the new database
```
\c gator
```

you should see a new prompt that looks like this:
```
gator+#
```

you can query the database to test
```
SELECT version();
```

you can type `exit` to leave psql shell


## install golang
### option 1(Linux/WSL/macOS): use the [webi installer](https://webinstall.dev/golang/)
```
curl -sS https://webi.sh/golang | sh
```

### option 2
Use the official [Golang installation instructions](https://go.dev/doc/install). On Windows, this means downloading and running a .msi installer package; the rest should be taken care of automatically.

After installing Golang, open a new shell session and run go version to make sure everything works. If it does, move on to the next steps

## install goose
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## to build the app
You can build with
```
go build ./cmd/gator
```

## install gator
The command to install the app is
```
go install ./cmd/gator
```

## run goose migrations

`cd` into `sql/schema` directory and run
```
goose postgres <connection_string> up
#example:
# goose postgres "postgres://wagslane:@localhost:5432/gator" up
```

the connection string is like this
`protocol://username:password@host:port/database`

For example:
* macOS (no password, your username): postgres://wagslane:@localhost:5432/gator
* linux (password from the setup, postgres user): postgres://postgres:postgres@localhost:5432/gator

test your connection string by running it with psql
`psql "postgres://wagslane:@localhost:5432/gator"`

## config file
create a config file in your home directory, `~/.gatorconfig.json`, with the following content:
Update with the proper connection string, need to add `?sslmode=disable`.

```json
{
    "db_url": "protocol://username:password@host:port/database?sslmode=disable"
}
```

## Examples on how to run the program

* To Reset the database, this is only needed if you want to start on a clean database
```#gator reset
Database was reset successfully
```

* To register a user, use the register command
```
#gator register kahya
User created sucdessfully:
 * ID: e44349e1-2687-49b8-aeff-fedce9400c69
 * Name: kahya
```

* To login a user, use the login command
```
#gator login kahya
User has been set to kahya
```

* To list users
`gator users`
```
* kahya (current)
```

* To add a feed for the current user
```
#gator addfeed "Hacker News RSS" "https://hnrss.org/newest"
Feed created succesfully
 * ID: 49fa2387-d860-4eed-8380-7759f5496339
 * Created At: 2026-07-12 22:32:49.741292 +0000 +0000
 * UpdatedAt: 2026-07-12 22:32:49.741292 +0000 +0000
 * Name: Hacker News RSS
 * Url: https://hnrss.org/newest
 * User: kahya

 * FeedName: Hacker News RSS
 * Username: kahya
========================================
```

* To list feeds
```
#gator feeds
 * ID: 49fa2387-d860-4eed-8380-7759f5496339
 * Created At: 2026-07-12 22:32:49.741292 +0000 +0000
 * UpdatedAt: 2026-07-12 22:32:49.741292 +0000 +0000
 * Name: Hacker News RSS
 * Url: https://hnrss.org/newest
 * User: kahya
=============================================
```

* Once feeds are registered, other users can also follow the same feed using the follow command. The user who adds a feed automatically follows it.
```
#gator follow https://hnrss.org/newest
Feed follow created succesfully
 * FeedName: Hacker News RSS
 * Username: holgith

======================================================
```

* To list what feeds are being follow by the logged user you can use the command following
```
#gator following
Feed follows for user holgith:
 * Name: Lanes Blog
 * Name: Hacker News RSS

===========================================================
```

* To unfollow a feed the `unfollow` command can be used
```
#gator unfollow https://www.wagslane.dev/index.xml
Lanes Blog unfollowed succesfully!
```

* The `agg` command would fetch the RSS feeds, parse them, and save them in the database. A duration interval is required as a paramter.
e.g. 1m, 10s, etc

```
#gator agg 1m
Collecting feeds every 1m0s...
Feed Hacker News RSS collected, 20 posts found
Created Post 0: Title: How to build a circular LCD clock
Created Post 1: Title: We taught our platform to learn its own pricing decisions
...
```

* The browse command receive an optional limit(default 2), it will list the saved posts for the user

```
#gator browse 3
Found 3 Posts for User holgith
===================================================
Sun Jul 12 from Hacker News RSS
--- How to build a circular LCD clock ---

<p>Article URL: <a href="https://blinry.org/lcd-clock/">https://blinry.org/lcd-clock/</a></p>
<p>Comments URL: <a href="https://news.ycombinator.com/item?id=48885613">https://news.ycombinator.com/item?id=48885613</a></p>
<p>Points: 1</p>
<p># Comments: 0</p>

Link: https://blinry.org/lcd-clock/
==============================================
Sun Jul 12 from Hacker News RSS
--- We taught our platform to learn its own pricing decisions ---

<p>Article URL: <a href="https://avriz.io/eng/paper">https://avriz.io/eng/paper</a></p>
<p>Comments URL: <a href="https://news.ycombinator.com/item?id=48885600">https://news.ycombinator.com/item?id=48885600</a></p>
<p>Points: 1</p>
<p># Comments: 0</p>

Link: https://avriz.io/eng/paper
==============================================
Sun Jul 12 from Hacker News RSS
--- LLMs and Shaders ---

<p>Article URL: <a href="https://amitp.blogspot.com/2026/07/llms-and-shaders.html">https://amitp.blogspot.com/2026/07/llms-and-shaders.html</a></p>
<p>Comments URL: <a href="https://news.ycombinator.com/item?id=48885572">https://news.ycombinator.com/item?id=48885572</a></p>
<p>Points: 1</p>
<p># Comments: 0</p>

Link: https://amitp.blogspot.com/2026/07/llms-and-shaders.html
==============================================
```
