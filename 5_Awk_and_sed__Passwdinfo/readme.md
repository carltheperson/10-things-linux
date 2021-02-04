# 5. Awk and sed - Passwdinfo

*Awk* and *sed* are more of the UNIX magic that I have always thought was really cool, though I never really understood what they were used for.
I often saw answers on Stack Overflow with people using them in crazy one-liners, but I always copy-pasted them without much thought.
Well, time to unravel the mystery.

I primarily used [this](https://www-users.york.ac.uk/~mijp1/teaching/2nd_year_Comp_Lab/guides/grep_awk_sed.pdf) paper to learn about them.
For the project, I wanted to create my own one-liner that shows information about the users on your system in a clear way.
I found just reading the */etc/passwd* a little too messy, so the project *passwdinfo*, displays the most important information in a neat table.
I found information about the */etc/passwd* file [here](https://www.cyberciti.biz/faq/understanding-etcpasswd-file-format/).

![](screenshot.png)
