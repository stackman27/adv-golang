package main

import (
	"fmt"
	"sort"
)


type Order struct {
	Id int
	Price float64
	Quantity int 
	IsBuy bool 
}

type OrderBook struct {
	BuyOrder []Order
	SellOrder []Order
}

func OrderbookMain() {
	fmt.Println("HELLO WORLD")

	newOrderBook := OrderBook{BuyOrder: nil, SellOrder: nil}

	buyOrder := Order {
		Id: 1,
		Price: 100,
		Quantity: 1,
		IsBuy: true,
	}

	sellOrder := Order {
		Id: 1,
		Price: 100,
		Quantity: 1,
		IsBuy: false,
	}

	sellOrder1 := Order {
		Id: 2,
		Price: 200,
		Quantity: 1,
		IsBuy: false,
	}


	sellOrder3 := Order {
		Id: 3,
		Price: 500,
		Quantity: 1,
		IsBuy: false,
	}


	updatedOrderbook := AddOrder(buyOrder, newOrderBook)
	updatedOrderbook = AddOrder(sellOrder, updatedOrderbook)
	updatedOrderbook = AddOrder(sellOrder1, updatedOrderbook)
	updatedOrderbook = AddOrder(sellOrder3, updatedOrderbook)
	fmt.Println("UPDATED ORDERBOOK: ", updatedOrderbook)

	updatedOrderbook = RemoveOrder(sellOrder1, updatedOrderbook)
	fmt.Println("UPDATED ORDERBOOK: ", updatedOrderbook)


	// match Order 
	newOrderBook = MatchOrder(updatedOrderbook) 
	fmt.Println("MATCHED ORDERBOOK: ", newOrderBook)
}

 

func AddOrder(order Order, existingOrderBook OrderBook) OrderBook{
	// adds a new order to the Orderbook  
	if order.IsBuy {
		// this is buy order 
		existingOrderBook.BuyOrder = append(existingOrderBook.BuyOrder, order) 
		// sort descending 
		sort.Slice(existingOrderBook.BuyOrder, func(i, j int) bool {
			return existingOrderBook.BuyOrder[i].Price > existingOrderBook.BuyOrder[j].Price
		})
		} else {
		// this is sell order 
		existingOrderBook.SellOrder = append(existingOrderBook.SellOrder, order) 
		// sort ascending
		sort.Slice(existingOrderBook.SellOrder, func(i, j int) bool {
			return existingOrderBook.SellOrder[i].Price < existingOrderBook.SellOrder[j].Price
		})
	}
	
	return existingOrderBook 
}


func RemoveOrder(orderToDelete Order, existingOrderBook OrderBook) OrderBook {
	if orderToDelete.IsBuy {
			// get the order Id 
			for i, order := range existingOrderBook.BuyOrder {
				if orderToDelete.Id == order.Id {
					// remove this buy order 
					existingOrderBook.BuyOrder = append(existingOrderBook.BuyOrder[:i], existingOrderBook.BuyOrder[i + 1:]...)
					 
				}
			}
	} else {
		for i, order := range existingOrderBook.SellOrder {
			if orderToDelete.Id == order.Id {
				// remove this sell order 
				existingOrderBook.SellOrder = append(existingOrderBook.SellOrder[:i], existingOrderBook.SellOrder[i + 1:]...)
				 
			}
		}
	}

	return existingOrderBook
}

func MatchOrder(existingOrderBook OrderBook)OrderBook{
	// if highest ask and lowest buy 
	// an order is matched is there is a buy Order >= sell order 
	var updatedOrderBook OrderBook
	existingbuyOrder := existingOrderBook.BuyOrder
	existingSellOrder := existingOrderBook.SellOrder
	
	for len(existingbuyOrder) > 0 && len(existingSellOrder) > 0 {
		orderSell := existingSellOrder [0]
		orderBuy := existingbuyOrder[0]

		if orderBuy.Price >= orderSell.Price {
			updatedOrderBook = RemoveOrder(orderBuy, existingOrderBook)
			updatedOrderBook = RemoveOrder(orderSell,  updatedOrderBook) 
			fmt.Printf("Matched %d @ %.2f\n", orderBuy.Quantity, orderBuy.Price)

			existingbuyOrder = updatedOrderBook.BuyOrder
			existingSellOrder = updatedOrderBook.SellOrder
		} else {
			break; 
		}
	} 

	return updatedOrderBook 
}

