# Market Module

This module takes in data, to produce a non-fungible item. This item will have a reference to a parent.

- The receiver of the payment of first purchase will the owner of the parent

## Create Item Data Needed

```golang
type Item struct {
  ItemID          UUID // unique identifier of the ticket
  OwnerName       string // owner of the item
  OwnerAddress    string // owner address
  ParentReference string || UUID // reference to parent in this case a event
  InitialPrice    sdk.Coins // original price of the item, if initialPrice is 0 then its a free event
  itemNumber      int // if the parent wants to make more than one
  Resale          bool
  MarkUpAllowed   int // amount of the current price (originalPrice || newPrice)
  ResaleCounter   int // amount of times it the item has been resold
  NewPrice        sdk.Coins // price that the item will be resold for
}
```

## Keeper

```golang
type Keeper struct {
	CKeeper bank.Keeper
	EKey    sdk.StoreKey // upcoming event key
	TKey    sdk.StoreKey // key for tickets that are generated for the people
	MKey    sdk.StoreKey // marketplace key for reselling
	UKey    sdk.StoreKey // store to keep an array of all the user tickets
	cdc     *codec.Codec
}
```

### Stores

- Market store: ket = event, data is tickets tickets that are up for sale
- Ticket store: key = event, data is tickets,

## Open Questions:

- Query tickets based on individuals, how?
- Dex: If we want to build it out as a dex then we can put bids and asks for tickets?
-

#### Query all tickets a individual holds, Options:

- 1. Have a auth module that has a store for each user that holds the tickets as data.
  - the primary tasks of this module is to only add and delete tickets from the users store.
  - key = sdk.AccAddress
- 2. Create a extra store to handle user tickets
  - key = user address, data is the array of all tickets owned then. open and closed events

### Uses

- Create Item using data from the parent
- Parent is original owner, they get the first full payment
- Parent gets percentage of resale
- Markup limit on each resale of the purchase price, not original price.
- NewPrice is derived from markup \* (originalPrice, after second then newPrice)

### User Flows

- As a user I want to get a free ticket to an event
- As a user I want to pay for one or more tickets to an event
- As a user I want to resell my ticket(s) on a marketpalce for more than I bought it for by specifying the price
- As a user I want to be able to send my ticket to someone else for free or for money
- As an event owner I want to receive a cut from the resale of my tickets only if its over the original price
- As a user I have to pay for the transaction fee (unless its a free event)
- As a user I want to see all the offered tickets for any given event on the marketplace
