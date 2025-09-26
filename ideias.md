## Entities

- **Player** → player data: base_rating, fatigue, morale, synergy, happiness, role, nationality, etc.  
- **Team** → esports team: name, tag, country, logo, list of players, aggregated stats.  
- **TeamPlayer** → association between Player ↔ Team (lineup, history, specific role).  
- **Match** → match between 2 teams: rounds, winner, round stats.  
- **Tournament** → tournament with multiple matches and final ranking.  
- **Facility** → team facilities: training center, psychologist, academy, strategy room, with effects on players/team.  
- **GameState / SaveInfo** → global save state: current day, player money, XP, active events, etc.  

---

## Systems

- **Match System** → controls match and round simulation, calculates winner based on stats, synergy, and randomness.  
- **Tournament System** → manages tournaments, brackets, matches, ranking, and rewards.  
- **Transfer System** → manages player hiring and transfers between teams, history, and salaries.  
- **Training System** → applies player progression considering morale, fatigue, facilities, and XP.  
- **Facility System** → manages facility construction and upgrades, applies effects to players and team.  
- **Team Management System** → manages lineup, player roles, team morale, synergy, and strategic adjustments.  
- **Finance System** → tracks player money, prize rewards, contracts, sponsorships, and upgrade costs.  
- **Event System** → handles random or scheduled events: injuries, suspensions, morale drops, sponsorship contracts.  
- **Player Progression System** → calculates player growth or decay based on performance, morale, and training.  