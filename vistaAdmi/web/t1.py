from mockito import when, mock, unstub

class Dog(object):
    def bark(self):
        return 'Wuff'


# either mock the class
when(Dog).bark().thenReturn('Miau!')
# now all instances have a different behavior
rex = Dog()
assert rex.bark() == 'Miau!'

# or mock a concrete instance
when(rex).bark().thenReturn('Grrrr')
assert rex.bark() == 'Grrrr'
# a different dog will still 'Miau!'
assert Dog().bark() == 'Miau!'

# be sure to call unstub() once in while
unstub()

##

from mockito import verify
# once again
rex = Dog()
when(rex).bark().thenReturn('Grrrr')

rex.bark()
rex.bark()

# `times` defaults to 1
verify(rex, times=2).bark()


##
#rex = Dog()
# this will fail because `Dog` has no method named `waggle`
when(rex).waggle().thenReturn('Nope')
# this will fail because `bark` does not take any arguments
when(rex).bark('Grrr').thenReturn('Nope')


# given this function
def bark(sound, post='!'):
    return sound + post

from mockito import kwargs
when(main).bark('Grrr', **kwargs).thenReturn('Nope')

# now this one will fail
bark('Grrr')  # because there are no keyword arguments used
# this one will fail because `then` does not match the function signature
bark('Grrr', then='!!')
# this one will go
bark('Grrr', post='?')

# there is also an args matcher
def add_tasks(*tasks, verbose=False):
    pass

from mockito import args
# If you omit the `thenReturn` it will just return `None`
when(main).add_tasks(*args)

add_tasks('task1', 'task2')  # will go
add_tasks()  # will fail
add_tasks('task1', verbose=True)  # will fail too

# On Python 3 you can also use `...`
when(main).add_tasks(...)
# when(main).add_tasks(Ellipsis) on Python 2

add_tasks('task1')  # will go
add_tasks(verbose=True)  # will go
add_tasks('task1', verbose=True)  # will go
add_tasks()  # will go
