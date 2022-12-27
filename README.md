**Requirements**
* UNIX based OS
* buf


**How To Run**
* run `make setup` to download dependencies
* run `make develop` to debug localy

If you want to run migration
* setup your db
* create `.env` file (you can copy example parameters feom `.env.example`)
* run `sudo make prepare-migration` to download `migrate`
* then run `make run-migration`