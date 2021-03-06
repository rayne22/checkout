package main

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"log"
	"net/http"
)

func main() {
	// This is your real test secret API key.
	stripe.Key = "sk_test_51G7PgLF8tXEBf9tJWq7hFIjcgndXfCNvg8LvdzQhYO5owePRIablydBvB2GMX4TfP2aDPyXEQidIOmqjOCiUByUr00MXBqqrcb"

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/create-checkout-session", createCheckoutSession)
	addr := "localhost:4242"
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// TODO: replace this with the `price` of the product you want to sell
				Price: stripe.String("{{PRICE_ID}}"),
				Quantity: stripe.Int64(1),
			},
		},
		PaymentMethodTypes: stripe.StringSlice([]string{
			"sepa_debit",
		}),
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL: stripe.String(domain + "/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
