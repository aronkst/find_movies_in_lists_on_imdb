# Find Movies in Lists on IMDb

This is a small project designed to find movies with a specific rating and by their amount of votes in a list of movies on the IMDb website.

# Parameters

- **1st:** IMDb List or IMDb URL

- **2nd:** Minimum value of the average movie rating in IMDb and Metascore (If cannot find the Metascore rating, the average rating will be the IMDb rating)

- **3rd:** Minimum amount of votes

# Example

`git clone https://github.com/aronkst/find_movies_in_lists_on_imdb`

`cd find_movies_in_lists_on_imdb`

`go get ./...`

`go build .`

`./find_movies_in_lists_on_imdb ls041125816 7.5 100000`

`./find_movies_in_lists_on_imdb "https://www.imdb.com/list/ls041125816/" 7.5 100000`
