# Try Eye

Eye is a process monitoring tool


requires ruby 2+
can also try on 1.9

### Run

    gem i bundler eye

    bundle

    eye help

    eye load config/thin.eye.rb -f

    # read the output, and you can now run:

    eye daemon
    
    eye load thin.eye.rb -f
    
    # this time it should boot (or complain about bundler, ruby versions, but hopefully not :])
