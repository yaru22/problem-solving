select
  country.continent,
  floor(avg(city.population))
from city
join country
on city.countrycode = country.code
group by country.continent
