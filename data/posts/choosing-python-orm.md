choosing-python-orm
Post
Choosing a Python ORM
2016-10-04





**Edit 2018:** This project is dead now, I will be able to make a new version in january but nothing certain.

When I started **FyPress** I was tired of *SQLAlchemy* and wanted to do something different and simple, my first idea was to do RAW SQL only. It was annoying and I was wasting a lot of time, so I started Fy_MySQL who was supposed to be a small MySQL ORM. At the time I really wanted to move forward on FyPress so I did not take the time and because of that Fy_MySQL is ugly. I wanted to migrate FyPress back to SQLAlchemy or Peewee but I also wanted to have fun and do something personal that’s why I started to do a “real” ORM: **fysql**.

**[SQLAlchemy](https://github.com/Fy-/fysql) and [Peewee](https://github.com/Fy-/fysql) are two great ORM and you should use them on your projects**, but if you want to learn and make mistakes I would also recommend you to try to do your own stuff if you have the time. FyPress could be a mini-CMS using Flask-SQLAlchemy, Flask-User, Flask-WTF, Flask-Security, etc… but it’s not, because I want something light and personnal. The downside is that I will spend more time developing already existing stuff but I don’t care, I will take the time. I’m not in a hurry, FyPress is not a corporate project. I don’t have the pretention to think I could make something better than all thoses libraries, but I can do something different and at least I will learn a lot.

Back to **fysql**, I wanted something simple with no “SQL like syntax”, Everything you need should be defined in your “Table” class.

I will migrate **FyPress** to fysql soon(ish) and I will use fysql for other personal projects so it will be maintained.

If you have some time please checkout **fysql**:

- [**fysql repository**](https://github.com/Fy-/fysql)
- [**fysql documentation**](https://github.com/Fy-/fysql)

Don’t hesitate to contact me if you have some ideas for fypress or fysql ^_^