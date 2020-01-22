package com.github.ezhydzetski.govsreactor;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.time.Duration;
import java.util.concurrent.TimeUnit;

@SpringBootApplication
public class BlockingUndertowApplication {
	public static void main(String[] args) {
		SpringApplication.run(BlockingUndertowApplication.class, args);
	}

	@RestController
	public static class TestController {
		@GetMapping(value = "/sleep", produces = MediaType.TEXT_PLAIN_VALUE)
		public String get(@RequestParam(required = false, defaultValue = "1s") String time) throws InterruptedException {
			Duration duration = Duration.parse("PT"+time);
			TimeUnit.NANOSECONDS.sleep(duration.toNanos());
			return "Hello after " + time;
		}
	}
}
