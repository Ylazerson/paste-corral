


<p align="center">  
    <img src=".img/paste-corral-logo.png" alt=""/>
</p>



Paste Corral crawls [pastebin.com](pastebin.com) to collect and store pastes. Paste Corral also provides a REST API endpoint for other developers to then easily consume pastebin data. 


<p align="center">  
    <font size="+3">Setup</font>
</p>


<p align="center">  
    <font size="+2">Prerequisite Steps</font>
</p>

 
1. Create an account on **Heroku**

2. Install **Heroku CLI** 
    - https://devcenter.heroku.com/articles/heroku-cli


<h4 style="text-align: center; color:#264653">
Setup Steps
</h4>

**Step 1:**
Fork and then clone this GitHub repo.


**Step 2:**
Run `heroku create`
- This creates a new empty application on Heroku, along with an associated empty Git repository. If you run this command from your app’s root directory, the empty Heroku Git repository is automatically set as a remote for your local repository.
- `git remote -v`
- Note, the file `Procfile` tells Heroku which command(s) to run to start your app.


**Step 3:**
Add a free Heroku Postgres Starter Tier dev database to your app:
- `heroku addons:create heroku-postgresql:hobby-dev`

Create a `.env` file.
- Note, this file is intentionally in `.gitignore`
- Add `PORT=8080` to the file

Show the `$DATABASE_URL` environment variable:
- `heroku config`
- Add that `DATABASE_URL` environment variable to the `.env` file.




zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz


#### Step 7:
- Run `heroku config` to get the Heroku app name
- Open that app on your Heroku dashboard: https://dashboard.heroku.com/apps
- Open the PostgreSQL add-on for that app.
- View Credentials on the Settings app.
- Create connection in **PostgreSQL Explorer** Extension

#### Step 8:
- Using pgAdmin run the `data/setup.sql` script.


#### Git it on up:
- Note this will be the general flow for working with Git now that we have Heroku remote as well.

```sh
go mod tidy
go mod vendor
go test

git status
git add --all
git commit -a  -m 'Initial launch'
git push heroku master

# Push to GitHub as well:
git push origin master
```

---

### Run your app on **Heroku**:

As a handy shortcut, you can open the website as follows:
- `heroku open`

View information about your running app:
- `heroku logs --tail`



--- 

## Run the app locally:

### Build package:

**Syntax**: `go build [-o output] [-i] [build flags] [packages]`

See https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
- `o` : write the resulting executable or object to the named output file or director
- `v` : prints the names of packages and files as they are processed


```sh
go build -o bin/chitchat -v .
```

- Start your app locally using the `heroku local` command.
    - This is installed as part of the Heroku CLI.
    - Just like Heroku, it examines the `Procfile` to determine what to run.

```sh
heroku local web
```

- Open http://localhost:8080 with your web browser. 
- To stop the app from running locally, go back to your terminal window and press `Ctrl+C` to exit.





---

## Setting Up A Custom Domain For Your Heroku-Hosted App

Note the glossary:
- https://devcenter.heroku.com/articles/custom-domains#domain-name-glossary

Approach taken here is based on:
- https://devcenter.heroku.com/articles/custom-domains
- https://medium.com/@ethanryan/setting-up-a-custom-domain-for-your-heroku-hosted-app-6c011e75aa3d

#### Step 1:
- Buy a custom domain name (I used name.com)
- Example used below is for domain `b7forum.com`

#### Step 2:

```sh
heroku domains:add www.b7forum.com

heroku domains:wait 'www.b7forum.com' 
```

Note the output; configure your app's DNS provider to point to the DNS Target:
- `integrative-kiwi-i8a0bariil1ahajauz37x4n6.herokudns.com`


#### Step 3: Add a custom root domain.

```sh
heroku domains:add b7forum.com

heroku domains:wait 'b7forum.com' 
```

Note the output, configure your app's DNS provider to point to the DNS Target:
- `still-sprout-wdcdz1tpzkyitakes8rs9x7z.herokudns.com`



#### Step 4:
- Add the DNS Records in name.com (or whatever site you bought your domain on).

![](docs/dns.png)


#### Step 5:
View existing domains:
- `heroku domains`

... Now wait a few minutes ...



---
NOT NEEDED YET:

#### Step ???:
 

- Set your config variables on Heroku.
- Note, `$PORT` is automatically set by Heroku on web dynos - so don't set that one.

`heroku config:set REPEAT=10`
