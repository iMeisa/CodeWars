def movie(card, ticket, perc)
    card_total = card
    card_prev = ticket
    ticket_total = 0

    movie_count = 0
    while card_total >= ticket_total
        discount = card_prev * perc
        card_prev = discount
        card_total += discount

        ticket_total += ticket

        movie_count += 1
    end

    return movie_count
end

puts movie(500, 15, 0.9)
puts movie(100, 10, 0.95)