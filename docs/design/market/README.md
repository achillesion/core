# Market Module

This module takes in data, to produce a non-fungible item. This item will have a reference to a parent.

- The receiver of the payment of first purchase will the owner of the parent

## Create Item Data Needed

```golang
type Item struct {
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

set key as event reference

### Uses

- Create Item using data from the parent
- Parent is original owner, they get the first full payment
- Parent gets percentage of resale
- Markup limit on each resale of the purchase price, not original price.
- NewPrice is derived from markup \* (originalPrice, after second then newPrice)

### User Flows

- As a user I want to get a free ticket to an event
- As a user I want to pay for a ticket to an event
- As a user I want to resell my ticket for more than I bought it for
- As an event owner I want to receive a cut from the resale of my tickets only if its over the original price
