package com.github.ezhydzetski.govsreactor;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

import java.time.Duration;

@SpringBootApplication
public class ReactiveNettyApplication {
	public static void main(String[] args) {
		SpringApplication.run(ReactiveNettyApplication.class, args);
	}

	@RestController
	public static class TestController {
		@GetMapping("/sleep")
		public Mono<String> get(@RequestParam(required = false, defaultValue = "1s") String time) {
			Duration duration = Duration.parse("PT"+time);
			return Mono.just("Hello after " + time)
			           .delaySubscription(duration);
		}
	}
}
