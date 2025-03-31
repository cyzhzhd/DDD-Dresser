package com.dresser.domain.prices;

import lombok.Getter;

import java.math.BigDecimal;
import java.util.Currency;

@Getter
public class Price {
    private final BigDecimal amount;
    private final Currency currency;
    
    private Price(BigDecimal amount, Currency currency) {
        if (amount == null) {
            throw new IllegalArgumentException("Amount cannot be null");
        }
        if (amount.compareTo(BigDecimal.ZERO) < 0) {
            throw new IllegalArgumentException("Amount cannot be negative");
        }
        if (currency == null) {
            throw new IllegalArgumentException("Currency cannot be null");
        }
        
        this.amount = amount;
        this.currency = currency;
    }
    
    public static Price of(BigDecimal amount, Currency currency) {
        return new Price(amount, currency);
    }
    
    public static Price of(BigDecimal amount, String currencyCode) {
        return new Price(amount, Currency.getInstance(currencyCode));
    }
    
    public static Price zero(Currency currency) {
        return new Price(BigDecimal.ZERO, currency);
    }
    
    public static Price zero(String currencyCode) {
        return new Price(BigDecimal.ZERO, Currency.getInstance(currencyCode));
    }
    
    public Price add(Price other) {
        if (!this.currency.equals(other.currency)) {
            throw new IllegalArgumentException("Cannot add prices with different currencies");
        }
        
        return new Price(this.amount.add(other.amount), this.currency);
    }
    
    public Price subtract(Price other) {
        if (!this.currency.equals(other.currency)) {
            throw new IllegalArgumentException("Cannot subtract prices with different currencies");
        }
        
        BigDecimal result = this.amount.subtract(other.amount);
        if (result.compareTo(BigDecimal.ZERO) < 0) {
            throw new IllegalArgumentException("Result cannot be negative");
        }
        
        return new Price(result, this.currency);
    }
    
    public Price multiply(int quantity) {
        if (quantity < 0) {
            throw new IllegalArgumentException("Quantity cannot be negative");
        }
        
        return new Price(this.amount.multiply(new BigDecimal(quantity)), this.currency);
    }
    
    @Override
    public String toString() {
        return amount.toString() + " " + currency.getCurrencyCode();
    }
} 