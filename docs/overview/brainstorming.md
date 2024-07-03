---
cover: >-
  https://images.unsplash.com/photo-1503551723145-6c040742065b-v2?crop=entropy&cs=srgb&fm=jpg&ixid=M3wxOTcwMjR8MHwxfHNlYXJjaHw3fHxicmFpbnN0b3JtfGVufDB8fHx8MTcyMDAxNzg1M3ww&ixlib=rb-4.0.3&q=85
coverY: 0
---

# 💡 Brainstorming

## User Inputs

* Capital
* Leverage
* Risk per trade (%)
* Number of trades for simulation
* Risk/Reward ratio (RR)

## Calculations

* Position size (number of lots)
* Stop loss ($, pips)
* Take profit ($, pips)
* Potential profit-and-loss (PnL)

## Outputs

* Position size (lot size)
* Graph of trade vs Potential PnL

## Rough Work

* 3:1 RR.
* 50 pips stop loss.
* 150 pips take profit.
* Number of trades = 100

Good day 100% WR $$= 15 \times 3 = \$45$$

Average day 67% WR $$= (15 \times 2) - 5 = 10 - 5 = \$25$$

Okay day 33% WR $$= 15 - (5  \times 2) = 15 -10 = \$5$$

Bad day 0% WR $$= - (5 \times 3) = -\$15$$

For $100 capital, PnL array = \[45, 25, 5, -15].

For $200 capital, PnL array = \[90, 50, 10, -30].

Plot all iterations on the same graph: days (x) and PnL (y).

