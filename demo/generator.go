package demo

// outline:
// [ ] add owner org
// - [ ] input:
// - - [ ] name string
// - - [ ] created/updated on time
// - [ ] output: owner org id
// [ ] add lab / storage org
// - [ ] input:
// - - [ ] name string
// - - [ ] created/updated on time
// - - [ ] owner org id
// - [ ] output:
// - - [ ] lab / storage org id
// [ ] add user
// - [ ] input:
// - - [ ] owner id
// - - [ ] (various value / enum inputs)
// - [ ] output:
// - - [ ] user id
// [ ] add plan
// - [ ] input:
// - - [ ] name string
// - - [ ] owner / lab storage org ids
// - [ ] output:
// - - [ ] plan id
// [ ] add protocol
// - [ ] input:
// - - [ ] owner id
// - - [ ] (various value / enum inputs)
// - [ ] output:
// - - [ ] protocol id
// [ ] add consent form
// - [ ] input:
// - - [ ] owner id
// - - [ ] title / body string
// - [ ] output:
// - - [ ] consent form id
// [ ] add protocol form
// - [ ] input:
// - - [ ] protocol id string (generated)
// - - [ ] protocol ids string array
// - - [ ] owner id
// - - [ ] title / body string
// - [ ] output:
// - - [ ] protocol form id
// [ ] add donor
// - [ ] input:
// - - [ ] owner id
// - - [ ] (various value / enum inputs)
// - - [ ] (not including consents / specimens)
// - [ ] output:
// - - [ ] donor id
// [ ] add consent
// - [ ] input:
// - - [ ] owner id
// - - [ ] donor id
// - - [ ] consent form id
// - - [ ] protocol id (non-generated)
// - - [ ] (various value / enum inputs)
// - - [ ] (not including specimens)
// - [ ] output:
// - - [ ] consent id
// [ ] add blood specimen
// - [ ] input:
// - - [ ] donor id
// - - [ ] consent id
// - - [ ] owner / lab / storage id
// - - [ ] protocol id
// - - [ ] (various value / enum inputs)
// - [ ] output:
// - - [ ] blood specimen id
// [ ] add test
// - [ ] input:
// - - [ ] owner id
// - - [ ] lab id
// - - [ ] specimens id
// - [ ] output:
// - - [ ] test id
// [ ] add result
// - [ ] input:
// - - [ ] owner id
// - - [ ] notes string
// - - [ ] test id
// - [ ] output:
// - - [ ] result id
