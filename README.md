# ard2rss
ard2rss is a small API that generates podcast RSS feeds for shows and collections of the [ARD Audiothek](https://www.ardaudiothek.de/). 

Since the ARD Audiothek does not provide RSS feeds directly, you can use this API to get your RSS feed and use it within your prefered podcast app.


ard2rss is currently hosted under https://ard2rss.mauamy.de.

# How to use
Go to the [ARD Audiothek](https://www.ardaudiothek.de/) and select a show or a collection you want an RSS feed for and copy the URL.  
Your URL should look like this:  
- https://www.ardaudiothek.de/sendung/<show_name>/<show_id>/
or
- https://www.ardaudiothek.de/sammlung/<collection_name>/<collection_id>

Now, to get your RSS feed simply replace `www.ardaudiothek.de` with `ard2rss.mauamy.de`.

For example:
https://www.ardaudiothek.de/sendung/levels-und-soundtracks/12642475/ --> https://ard2rss.mauamy.de/sendung/levels-und-soundtracks/12642475/

You can add this URL to your podcatcher/podcast app.





