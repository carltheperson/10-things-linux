# 9. Systemd services - Createservice

Whenever I try to set up a database on a Linux machine I have been confused about how to configure the *systemd* service.
I have also been in situations where I needed to create a service from a binary but always struggled.
The struggle ends now.

As with any new subject, it’s always a good idea to read the Wikipedia page, so that is where I started.
Surprisingly, I learned that *systemd* is a quite controversial piece of software, but I still wanted to learn it and judge it for myself.

For understanding the basics of *systemd* I read [this](https://www.digitalocean.com/community/tutorials/how-to-use-systemctl-to-manage-systemd-services-and-units) and for understanding how to create a new service I read [this](https://www.tecmint.com/create-new-service-units-in-systemd/).

For the project, I made *createservice*, which allows you to make a *systemd* service from any executable that will automatically start up on boot.
Here I test it out on Prometheus:

![](screenshot.png)
