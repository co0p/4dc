Shareit
======

A platform for sharing your personal items with your friends. It allows you to see what you friends have lying around, ask to borrow it. Everything is neatly tracked using this mobile web-app. 

functionalities
----------------
you can add items
borrow items
return items
view the item catalog
get details for each item

tech
---
modern mobile first webapp, tailwindcss and sample express backend deployed as one docker container to heroku. 

Quickstart (Catalog Browsing increment)
--------------------------------------
Run the catalog browsing demo locally:

```sh
# from repo root
npm --prefix examples/shareit run start:catalog
```

Then verify:

```sh
# API returns items
curl http://localhost:3000/api/items | python3 -m json.tool

# Open the SPA in your browser
open http://localhost:3000
```

Notes:
- Vue 3 and Tailwind are loaded via CDN for speed.
- Data is served from an in-memory repository and resets on restart.