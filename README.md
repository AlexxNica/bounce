# What is Bounce?

Bounce is an experimental authentication prototype. Bounce uses Ricochet as a means to offer authenticated pseudonyms!

This has several advantages over traditional user management solutions:

* **Anonymity** - Using the power of Ricochet and Tor, the user can disassociate themselves from the account, without having to setup an anonymous email account - which is often fraught with difficulties.
* **No Passwords to remember** - Authentication is provided through the authentication properties of Tor Hidden Services (through Ricochet). Of course, there is nothing to stop someone extending Bounce to also ask for another authentication factor! Which brings us to...
* **Decentralization** - Anyone can run a Bounce server, allowing users to pick and choose which entities get to know which of their ricochet addresses. Sites do not have to rely 3rd party providers to allow a consistent user experience across different sites.

# How does Bounce Work?

Bounce is fairly simple. First the user is asked for a ricochet id. Bounce then constructs an encrypted token and sends a message to the ricochet id.

The message contains a link which the user can click on to authenticate themselves to the site. The link contains an encrypted token which can be verified by the bounce service.

Bounce is based on <a href="https://github.com/s-rah/go-ricochet"/>go-ricochet</a>
