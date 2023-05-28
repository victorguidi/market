<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	const stocksData = writable({});
	const selectedStock = writable('');
	const options = writable({});
	const chart = writable(null);
	const stockFundamentals = writable({});
	const news = writable([]);

	$: stocksData.subscribe((data) => {
		options.set({
			chart: {
				type: 'candlestick',
				background: '#141823',
				roundedCorners: true
			},
			// TODO: Change the color of the info box
			series: [
				{
					name: $selectedStock,
					data: data[$selectedStock]?.values.map((v) => {
						return {
							x: new Date(v.datetime),
							y: [v.open, v.high, v.low, v.close]
						};
					}),
					labels: data[$selectedStock]?.values.map((v) => {
						return v.datetime;
					})
				}
			]
		});
	});

	function updateOptions(stock: string) {
		const chart = new ApexCharts(document.querySelector('#chart'), $options);
		chart.render();
		selectedStock.set(stock);
		options.update((o) => {
			return {
				...o,
				series: [
					{
						name: stock,
						data: $stocksData[stock]?.values.map((v) => {
							return {
								x: new Date(v.datetime),
								y: [v.open, v.high, v.low, v.close]
							};
						})
					}
				]
			};
		});
		$chart.updateOptions($options);

		fetch('http://localhost:8080/api/v1/stocks/' + $selectedStock)
			.then((res) => res.json())
			.then((data) => {
				stockFundamentals.set(data);
			});
	}

	const USDollar = new Intl.NumberFormat('en-US', {
		style: 'currency',
		currency: 'USD'
	});

	onMount(async () => {
		await fetch('http://localhost:8080/api/v1/users/stocks/daily/1')
			.then((res) => res.json())
			.then((data) => {
				for (let d in data) {
					data[d].values = data[d].values.reverse();
				}

				selectedStock.set(Object.keys(data)[0]);
				stocksData.set(data);
			});

		await fetch('http://localhost:8080/api/v1/stocks/' + $selectedStock)
			.then((res) => res.json())
			.then((data) => {
				console.log(data);
				stockFundamentals.set(data);
			});

		await fetch('http://localhost:8080/api/v1/rss/')
			.then((res) => res.json())
			.then((data) => {
				news.set(data);
			});

		chart.set(new ApexCharts(document.querySelector('#chart'), $options));
		$chart.render();

		const pie = new ApexCharts(document.querySelector('#pie'), {
			chart: {
				type: 'pie'
			},
			series: [55, 40, 5],
			labels: ['Buy', 'Sell', 'Hold'],
			legend: {
				position: 'bottom'
			},
			colors: ['#022E1F', '#FF0000', '#FFA500'],
			stroke: {
				show: false
			}
		});
		pie.render();
	});
</script>

<div class="flex flex-row justify-between w-screen h-screen p-4 bg-zinc-950 text-white">
	<div class="flex flex-col w-2/12 h-full rounded-md border-r border-zinc-900 p-4 bg-zinc-900">
		<h1 class="text-center mb-4 text-xl">Welcome Victor Guidi</h1>
		{#each Object.keys($stocksData) as stock}
			<button
				style={stock === $selectedStock ? 'background-color: #1abc9c' : ''}
				class="flex flex-col p-2 h-12 w-full rounded-md items-center justify-center mb-3"
				on:click={() => updateOptions(stock)}>{stock}</button
			>
		{/each}
		<button class="flex flex-col p-2 h-12 w-full border rounded-md items-center justify-center mb-3"
			>ADD</button
		>
	</div>
	<div class="flex flex-col w-8/12 h-full">
		<div id="chart" class="w-full m-6 h-1/2" />
		<div class="flex flex-col w-full p-4 h-1/2">
			<div class="flex w-fulljustify-between">
				<div class="rounded-md ml-2 w-full p-2 bg-teal-900">
					<h1 class="mb-2">Market Cap</h1>
					<div>{USDollar.format($stockFundamentals.MarketCapitalization)}</div>
				</div>
				<div class="rounded-md ml-2 w-full p-2 bg-teal-900">
					<h1 class="mb-2">Book Value</h1>
					<div>{USDollar.format($stockFundamentals.BookValue)}</div>
				</div>
				<div class="rounded-md ml-2 w-full p-2 bg-teal-900">
					<h1 class="mb-2">EBITDA</h1>
					<div>{$stockFundamentals.EBITDA}</div>
				</div>
				<div class="rounded-md ml-2 w-full p-2 bg-teal-900">
					<h1 class="mb-2">PERatio</h1>
					<div>{$stockFundamentals.PERatio}</div>
				</div>
				<div class="rounded-md ml-2 w-full p-2 bg-teal-900">
					<h1 class="mb-2">EPS</h1>
					<div>{$stockFundamentals.EPS}</div>
				</div>
				<div class="rounded-md ml-2 w-full p-2 bg-teal-900">
					<h1 class="mb-2">Dividends</h1>
					<div>{$stockFundamentals.DividendPerShare}</div>
				</div>
			</div>
		</div>
	</div>
	<div class="flex flex-col w-3/12 bg-zinc-900 p-2">
		<h1 class="text-xl">Recommendation</h1>
		<div class="flex flex-col h-1/2 justify-center rounded-md">
			<div id="pie" />
		</div>
		<div
			class="flex flex-col h-1/2 overflow-auto pr-3 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:'none'] [scrollbar-width:'none']"
		>
			<h1 class="mb-2 text-xl">News</h1>
			{#each $news as n}
				<div class="flex flex-row">
					<div class="flex flex-col p-2 h-max-26 w-full mb-3 rounded-md bg-neutral-800">
						<button class="flex flex-col items-start text-left w-full h-full"
							><a href={n?.link}>
								<div class="font-sans text-sm mb-2">{n.title}</div>
								<div class="text-sm">{n.published}</div>
							</a>
						</button>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
