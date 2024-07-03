---
cover: >-
  https://images.unsplash.com/photo-1600132806370-bf17e65e942f?crop=entropy&cs=srgb&fm=jpg&ixid=M3wxOTcwMjR8MHwxfHNlYXJjaHwzfHxibHVlcHJpbnR8ZW58MHx8fHwxNzIwMDE4MDA5fDA&ixlib=rb-4.0.3&q=85
coverY: 0
---

# 👷 Requirements Engineering

## User Requirements

* Risky should visualise the performance of users' trading strategies via an automated stochastic simulation.
* Risky should provide information about the lot size, stop loss pips and take profit pips that a user can take based on their risk tolerance level.

## System Requirements

### General

* The application should **collect and validate** user inputs.
* The application should **correctly calculate values** from user inputs.
* The application should be able to **visualise** the trading simulation.
* The application should not need a guide to be useful.

#### What are the user inputs?

* Capital
* Leverage
* Risk per trade (%)
* Number of trades for simulation
* Risk/Reward ratio (RR)

#### What are the values to be calculated?

* Position size (number of lots)
* Stop loss ($, pips)
* Take profit ($, pips)
* Potential profit-and-loss (PnL)

#### What is to be visualised?

* Days (x) and PnL (y) base plotting area.
* Graphs of overall simulations.

### Simulation

* The number of the **entire simulation iterations should be 5** to avoid noise.
* Win rate, WR, (%) should be broken into the following **4 levels**:

<table data-full-width="false"><thead><tr><th>Day Meter</th><th>Win Rate (WR)</th></tr></thead><tbody><tr><td>Good</td><td>100%</td></tr><tr><td>Average</td><td>67%</td></tr><tr><td>Okay</td><td>33%</td></tr><tr><td>Bad</td><td>0%</td></tr></tbody></table>

{% hint style="info" %}
The profit-and-loss (PnL) array should be derived from WR levels.
{% endhint %}

* The graph should be a **stochastic simulation** of the Day Meter WRs.
* The system should calculate the **average PnL** from the simulation iterations.
* The system should calculate the **net PnL** from each simulation iteration.
