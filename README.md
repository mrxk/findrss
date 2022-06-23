# findrss

This project creates a docker image that, when run, hosts a web server
that presents a single html page. There is one input dialog on that page.
When a user adds a URL to that dialog and clicks the button, the server
will parse the HTML returned for that URL and show all of the &lt;link&gt;
elements. This is intended to aid the user in discovering RSS and Atom links
for adding to a feed reeder.
