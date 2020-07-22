


<p align="center">  
    <img src=".img/paste-corral-logo.png" alt=""/>
</p>



Paste Corral crawls [pastebin.com](pastebin.com) to collect, clean, and store PasteBin posts. Paste Corral concurrently provides a REST API endpoint so developers can easily consume cleaned PasteBin data to perform analytics. 

See [www.pastecorral.com](www.pastecorral.com) for a live version of Paste Corral.
- At the moment it only supports a simple GET request.
- You can test it using `curl -i -X GET www.pastecorral.com`



## Setup

**Step 1:**
Fork and then clone this GitHub repo.


**Step 2:**
Create an account on **Heroku**


**Step 3:**
Install **Heroku CLI** 
- https://devcenter.heroku.com/articles/heroku-cli


**Step 4:**
Run `heroku create`
- This creates a new empty application on Heroku, along with an associated empty Git repository. 
- Run this command from your app’s root directory, so the empty Heroku Git repository is automatically set as a remote for your local repository.


**Step 5:**
Add a free Heroku Postgres Starter Tier dev database to your app:
- `heroku addons:create heroku-postgresql:hobby-dev`

Create a `.env` file.
- Note, this file is intentionally in `.gitignore`
- Add `PORT=8080` to the file

Show the `$DATABASE_URL` environment variable:
- `heroku config`
- Add that `DATABASE_URL` environment variable to the `.env` file.


**Step 6:**
Connect to the Heroku PostgreSQL instance and run the `data/setup.sql` script.

You can view your credentials using the `heroku config` command.

You can connect using any PostgreSQL admin tool. If you're using VSCode, the *PostgreSQL Explorer* extension works great. 



---

## General Notes

If you make any code changes, remember to:
- Commit and push the changes to your GitHub repo
- Then push to Heroku as well: `git push heroku master`   

To view information about your running Heroku app:
- `heroku logs --tail`

To open your Heroku app (in this case a REST API endpoint):
- `heroku open`

