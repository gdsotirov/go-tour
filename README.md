# Go-Tour

The exercises from my [Go Language Tour](https://go.dev/tour/) after
I first took it in beginning of August, 2014 and then kept refreshing
my knowledge whenever necessary.

## Most important about Go language

This is a summary of the most important points to consider with Go language
when taking the tour and afterwards when using the language.

 - programs are made of _packages_, programs start running in package `main`
   (see [Packages](https://go.dev/tour/basics/1));
 - a name is _exported_ from a package and can be refereed to only if it
   begins with a capital letter (see
   [Exported names](https://go.dev/tour/basics/3));
 - functions could return _multiple results_, including named return values
   (see [Multiple results](https://go.dev/tour/basics/6),
   [Named return values](https://go.dev/tour/basics/7));
 - variables without initializers are given _zero values_ (see
   [Zero values](https://go.dev/tour/basics/12));
 - type conversions are always _explicit_ (see
   [Type conversions](https://go.dev/tour/basics/13]));
 - there is only one _looping construct_ - `for` (see
   [For](https://go.dev/tour/flowcontrol/1),
   [For is Go's "while"](https://go.dev/tour/flowcontrol/3));
 - `switch` statement has implicit `break` after each case (see
   [Switch](https://go.dev/tour/flowcontrol/9));
 - `defer` statement for _deferring the execution_ of a function until
   the surrounding function returns (see
   [Defer](https://go.dev/tour/flowcontrol/12));
 - there are _pointers_, but no pointer arithmetic (see
   [Pointers](https://go.dev/tour/moretypes/1));
 - _slices_ are "dynamically-sized, flexible views into the elements of
   an array" (see [Slices](https://go.dev/tour/moretypes/7));
 - functions are _values_ and may be _closures_ as well (see
   [Function values](https://go.dev/tour/moretypes/24),
   [Function closures](https://go.dev/tour/moretypes/25));
 - no _classes_, but _methods on types_, which are functions with a special
   _receiver_ argument (see [Methods](https://go.dev/tour/methods/1));
 - _interface type_ is defined as a set of method signatures and a value of
   such type can hold any value that implements those methods (see
   [Interfaces](https://go.dev/tour/methods/9));
 - _error state_ is expressed with `error` values (see
   [Errors](https://go.dev/tour/methods/19));
 - _type parameters_ and _generic types_ are the way Go implements _generics_
   since version 1.18 (see
   [Type parameters](https://go.dev/tour/generics/1),
   [Generic types](https://go.dev/tour/generics/2));
 - _goroutines_ are a lightweight threads managed by the Go runtime (see
  [Goroutines](https://go.dev/tour/concurrency/1)), which communicate and
  synchronize through _channels_ (see
  [Channels](https://go.dev/tour/concurrency/2)) as sends and receives on
  channels block until the other side is ready;

## Usage

To build the exercises do the following:

 1. Clone the repository (e.g. in your home directory):

```
git clone --recurse-submodules https://github.com/gdsotirov/go-tour.git
cd go-tour
```

 2. Build all exercises:

```
make --jobs
```

 3. Run exercises (e.g. `ex_mtp_23_wc`):

```
./ex_mtp_23_wc
```

